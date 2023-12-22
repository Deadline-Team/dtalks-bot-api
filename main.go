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
	"context"
	"encoding/json"
	"fmt"
	"github.com/deadline-team/dtalks-bot-api/model"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/service"
	"github.com/deadline-team/dtalks-bot-api/util"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type BotAPI interface {
	service.AttachmentService
	service.CalendarEventService
	service.ConversationService
	service.LabelService
	service.MessageService
	service.ReactionService
	service.TaskService
	service.UserService

	// GetEventChannel
	// Метод для получения канала в который складываются все события с сервера
	GetEventChannel() chan model.Event
}

type botAPI struct {
	model.BotBaseParam

	service.AttachmentService
	service.CalendarEventService
	service.ConversationService
	service.LabelService
	service.MessageService
	service.ReactionService
	service.TaskService
	service.UserService

	tokenInfo model.TokenInfo
	channel   chan model.Event
	wsClient  *util.WebsocketClient
}

func New(host string, apiKey string, secure bool) (BotAPI, error) {
	tokenInfo, err := util.ParseToken(apiKey)
	if err != nil {
		return nil, err
	}

	botBaseParam := model.BotBaseParam{
		Host:   host,
		ApiKey: apiKey,
		Secure: secure,
	}

	return &botAPI{
		BotBaseParam:         botBaseParam,
		AttachmentService:    service.NewAttachmentService(botBaseParam),
		CalendarEventService: service.NewCalendarEventService(botBaseParam),
		ConversationService:  service.NewConversationService(botBaseParam),
		LabelService:         service.NewLabelService(botBaseParam),
		MessageService:       service.NewMessageService(botBaseParam),
		ReactionService:      service.NewReactionService(botBaseParam),
		TaskService:          service.NewTaskService(botBaseParam),
		UserService:          service.NewUserService(botBaseParam),
		tokenInfo:            tokenInfo,
		channel:              make(chan model.Event),
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

func (client *botAPI) createRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	schema := "http"
	if client.Secure {
		schema = "https"
	}

	request, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s://%s%s", schema, client.Host, path), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	request.Header.Set("Content-Type", "application/json")
	return request, err
}
