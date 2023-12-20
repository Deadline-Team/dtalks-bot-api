package dtalks_bot_api

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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deadline-team/dtalks-bot-api/model"
	attachmentModel "github.com/deadline-team/dtalks-bot-api/model/attachment"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/util"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type BotAPI interface {
	GetEventChannel() chan model.Event
	GetConversations() ([]conversationModel.Conversation, error)
	SendMessage(conversationId string, text string, attachments []*attachmentModel.Attachment) (*conversationModel.Message, error)
	CreateAttachment(fileName string, data []byte) (*attachmentModel.Attachment, error)
}

var httpClient = &http.Client{Timeout: time.Second * 30}

type botAPI struct {
	Host   string
	ApiKey string
	Secure bool

	tokenInfo model.TokenInfo
	channel   chan model.Event
	wsClient  *util.WebsocketClient
}

func New(host string, apiKey string, secure bool) (BotAPI, error) {
	tokenInfo, err := util.ParseToken(apiKey)
	if err != nil {
		return nil, err
	}

	return &botAPI{
		Host:      host,
		ApiKey:    apiKey,
		Secure:    secure,
		tokenInfo: tokenInfo,
		channel:   make(chan model.Event),
	}, nil
}

func (client *botAPI) GetEventChannel() chan model.Event {
	schema := "ws"
	if client.Secure {
		schema = "wss"
	}
	wsConfig := util.WebsocketConfig{
		Url:             fmt.Sprintf("%s://%s/api/notification/ws?access_token=%s", schema, client.Host, client.ApiKey),
		OnConnect:       client.onConnect,
		OnMessage:       client.onMessage,
		PingPeriod:      15 * time.Second,
		ReconnectPeriod: 15 * time.Second,
		IsAutoReconnect: true,
	}

	if client.wsClient != nil {
		client.wsClient.Terminate()
	}
	client.wsClient = util.NewWebsocketClient(wsConfig)
	go client.wsClient.Start()

	return client.channel
}

func (client *botAPI) GetConversations() ([]conversationModel.Conversation, error) {
	request, err := client.createRequest(http.MethodGet, "/api/conversation/conversations", nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	var conversationPage model.Page[conversationModel.Conversation]
	if err := json.NewDecoder(response.Body).Decode(&conversationPage); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return conversationPage.Content, nil
}

func (client *botAPI) SendMessage(conversationId string, text string, attachments []*attachmentModel.Attachment) (*conversationModel.Message, error) {
	data, err := json.Marshal(conversationModel.Message{Text: text, Attachments: attachments})
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(http.MethodPost, fmt.Sprintf("/api/conversation/conversations/%s/messages", conversationId), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}
	var message conversationModel.Message
	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &message, err
}

func (client *botAPI) CreateAttachment(fileName string, data []byte) (*attachmentModel.Attachment, error) {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)
	fw, err := bw.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	_, err = fw.Write(data)
	if err != nil {
		return nil, err
	}
	if err = bw.Close(); err != nil {
		return nil, err
	}

	request, err := client.createRequest(http.MethodPost, "/api/attachment/attachments", buf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", bw.FormDataContentType())

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	var attachment attachmentModel.Attachment
	if err := json.NewDecoder(response.Body).Decode(&attachment); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &attachment, err
}

func (client *botAPI) onConnect() {
	log.Printf("Connected like %s (%s)", client.tokenInfo.GetFullName(), client.tokenInfo.Username)
}

func (client *botAPI) onMessage(msg []byte) {
	var event model.Event
	if err := json.Unmarshal(msg, &event); err != nil {
		fmt.Println(err)
	}

	if event.Type == "NEW_MESSAGE_IN_CONVERSATION" {
		var message conversationModel.Message
		data, err := json.Marshal(event.Payload)
		if err != nil {
			fmt.Println(err)
		}
		if err := json.Unmarshal(data, &message); err != nil {
			fmt.Println(err)
		}

		if strings.HasPrefix(message.Text, "/") {
			event.UserId = ""
			event.Type = "Command"
			event.Payload, _ = strings.CutPrefix(message.Text, "/")
			client.channel <- event
		} else {
			event.UserId = ""
			event.Type = "Message"
			client.channel <- event
		}
	} else {
		client.channel <- event
	}
}

func (client *botAPI) createRequest(method string, path string, body io.Reader) (*http.Request, error) {
	schema := "http"
	if client.Secure {
		schema = "https"
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s://%s%s", schema, client.Host, path), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	request.Header.Set("Content-Type", "application/json")
	return request, err
}
