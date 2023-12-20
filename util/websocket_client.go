package util

/*
 * Copyright © 2023, "DEADLINE TEAM" LLC
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are not permitted.
 *
 * THIS SOFTWARE IS PROVIDED BY "DEADLINE TEAM" LLC "AS IS" AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL "DEADLINE TEAM" LLC BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * No reproductions or distributions of this code is permitted without
 * written permission from "DEADLINE TEAM" LLC.
 * Do not reverse engineer or modify this code.
 *
 * © "DEADLINE TEAM" LLC, All rights reserved.
 */

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var (
	cxt    context.Context
	cancel context.CancelFunc
)

func init() {
	cxt, cancel = context.WithCancel(context.Background())
}

type WebsocketConfig struct {
	Url              string
	ProxyUrl         string
	ReqHeaders       map[string][]string
	PingPeriod       time.Duration
	IsAutoReconnect  bool
	OnConnect        func()
	OnMessage        func([]byte)
	OnClose          func(int, string) error
	OnError          func(err error)
	ReadDeadlineTime time.Duration
	ReconnectPeriod  time.Duration
}

type WebsocketClient struct {
	WebsocketConfig
	dialer             *websocket.Dialer
	conn               *websocket.Conn
	writeBuffer        chan []byte
	closeMessageBuffer chan []byte
	reconnectLock      *sync.Mutex
	waiter             *sync.WaitGroup
	open               chan bool
}

func NewWebsocketClient(config WebsocketConfig) *WebsocketClient {
	client := &WebsocketClient{WebsocketConfig: config}
	client.dialer = &websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  30 * time.Second,
		EnableCompression: true,
	}
	if client.ProxyUrl != "" {
		proxy, err := url.Parse(client.ProxyUrl)
		if err != nil {
			logrus.Panic(err)
		} else {
			client.dialer.Proxy = http.ProxyURL(proxy)
		}
	}
	if config.PingPeriod == 0 {
		client.PingPeriod = 15 * time.Second
	}
	if config.ReadDeadlineTime == 0 {
		client.ReadDeadlineTime = 2 * client.PingPeriod
	}
	client.reconnectLock = new(sync.Mutex)
	client.waiter = &sync.WaitGroup{}
	return client
}

func (client *WebsocketClient) dial() error {

	conn, resp, err := client.dialer.Dial(client.Url, client.ReqHeaders)
	if err != nil {
		if resp != nil {
			dumpData, _ := httputil.DumpResponse(resp, true)
			err_ := errors.New(string(dumpData))
			logrus.Error(err_)
		}
		logrus.Errorf("websocket-client dial %s fail", client.Url)
		return err
	}
	client.conn = conn
	_ = client.conn.SetReadDeadline(time.Now().Add(client.ReadDeadlineTime))
	if client.OnConnect != nil {
		client.OnConnect()
	}
	if client.OnClose != nil {
		client.conn.SetCloseHandler(client.OnClose)
	}

	client.conn.SetPongHandler(func(appData string) error {
		_ = client.conn.SetReadDeadline(time.Now().Add(client.ReadDeadlineTime))
		return nil
	})

	return nil
}

func (client *WebsocketClient) initBufferChan() {
	client.open = make(chan bool)
	client.writeBuffer = make(chan []byte, 10)
	client.closeMessageBuffer = make(chan []byte, 10)
}

func (client *WebsocketClient) Start() {
	// buffer channel reset
	client.initBufferChan()
	// dial
	err := client.dial()
	if err != nil {
		logrus.Error("websocket-client start error:", err)
		if client.IsAutoReconnect {
			client.reconnect(10)
		}
	}
	// start read goroutine and write goroutine
	for {
		client.waiter.Add(2)
		go client.write()
		go client.read()
		client.waiter.Wait()
		close(client.open)
		if client.IsAutoReconnect {
			client.reconnect(40)
		} else {
			logrus.Info("websocket-client closed. bye")
			return
		}
	}
}

func (client *WebsocketClient) write() {
	var err error
	ctxW, cancelW := context.WithCancel(cxt)
	pingTicker := time.NewTicker(client.PingPeriod)
	defer func() {
		pingTicker.Stop()
		cancelW()
		client.waiter.Done()
	}()

	for {
		select {
		case <-ctxW.Done():
			logrus.Warn("websocket-client connect closing, exit writing progress...")
			return
		case d := <-client.writeBuffer:
			err = client.conn.WriteMessage(websocket.TextMessage, d)
		case d := <-client.closeMessageBuffer:
			err = client.conn.WriteMessage(websocket.CloseMessage, d)
		case <-pingTicker.C:
			err = client.conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		default:
			err = nil
		}
		if err != nil {
			logrus.Errorf("write error: %v", err)
			return
		}
	}
}

func (client *WebsocketClient) read() {
	ctxR, cancelR := context.WithCancel(cxt)
	defer func() {
		_ = client.conn.Close()
		cancelR()
		client.waiter.Done()
	}()
	for {
		select {
		case <-ctxR.Done():
			logrus.Warn("websocket-client connect closing, exit receiving progress...")
			return
		default:
			_ = client.conn.SetReadDeadline(time.Now().Add(client.ReadDeadlineTime))
			_, msg, err := client.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					logrus.Errorf("read error: %v", err)
				}
				return
			}
			if client.OnMessage != nil {
				client.OnMessage(msg)
			}
		}
	}
}

func (client *WebsocketClient) Send(msg []byte) {
	client.writeBuffer <- msg
}

func (client *WebsocketClient) SendClose() {
	client.closeMessageBuffer <- websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
}

func (client *WebsocketClient) Terminate() {

	client.reconnectLock.Lock()
	client.IsAutoReconnect = false
	defer client.reconnectLock.Unlock()
	cancel()
	<-client.open
}

func (client *WebsocketClient) reconnect(retry int) {
	client.reconnectLock.Lock()
	defer client.reconnectLock.Unlock()
	var err error
	client.initBufferChan()
	for i := 0; i < retry; i++ {
		err = client.dial()
		if err != nil {
			logrus.Errorf("websocket-client reconnect fail: %d", i)
		} else {
			break
		}
		time.Sleep(client.ReconnectPeriod)
	}
	if err != nil {
		logrus.Error("websocket-client retry reconnect fail. exiting....")
		client.Terminate()
	}
}
