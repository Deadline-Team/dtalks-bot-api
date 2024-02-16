package dtalks_bot_api

/*
 * Copyright © 2023 - 2024, "DEADLINE TEAM" LLC
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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deadline-team/dtalks-bot-api/model"
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

	// RegisterCommand
	// Метод для регистрации команд бота
	RegisterCommand(ctx context.Context, commandsAndDescriptions map[string]string) error

	// SendEvent
	// Метод для отправки событий на сервер через WebSocket
	SendEvent(ctx context.Context, event model.Event) error
}

type botAPI struct {
	model.BotBaseParam

	service.AttachmentService
	service.CalendarEventService
	service.ConversationService
	service.LabelService
	service.LinkService
	service.MessageService
	service.ReactionService
	service.TaskService
	service.UserService

	tokenInfo  model.TokenInfo
	channel    chan model.Event
	wsClient   *util.WebsocketClient
	httpClient *http.Client
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
		LinkService:          service.NewLinkService(botBaseParam),
		MessageService:       service.NewMessageService(botBaseParam),
		ReactionService:      service.NewReactionService(botBaseParam),
		TaskService:          service.NewTaskService(botBaseParam),
		UserService:          service.NewUserService(botBaseParam),
		tokenInfo:            tokenInfo,
		channel:              make(chan model.Event),
		httpClient:           &http.Client{Timeout: time.Second * 30},
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

func (client *botAPI) RegisterCommand(ctx context.Context, commandsAndDescriptions map[string]string) error {
	data, err := json.Marshal(&commandsAndDescriptions)
	if err != nil {
		return err
	}

	request, err := util.CreateHttpRequest(ctx, client.BotBaseParam, http.MethodPut, fmt.Sprintf("/api/bot/applications/user/%s/command", client.tokenInfo.UserId), bytes.NewReader(data))
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

func (client *botAPI) SendEvent(ctx context.Context, event model.Event) error {
	event.UserId = client.tokenInfo.UserId
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	client.wsClient.Send(data)
	return nil
}

func (client *botAPI) onConnect() {
	log.Printf("Connected like %s (%s)", client.tokenInfo.GetFullName(), client.tokenInfo.Username)
}

func (client *botAPI) onMessage(msg []byte) {
	var event model.Event
	if err := json.Unmarshal(msg, &event); err != nil {
		fmt.Println(err)
	}

	if strings.HasPrefix(event.Type, "NEW_MESSAGE") {
		message, err := util.ParseMessage(event)
		if err != nil {
			fmt.Println(err)
		}

		event.ConversationId = message.Meta["conversationId"].(string)
		if event.Type == "NEW_MESSAGE_IN_CONVERSATION" {
			if err := client.ReadMessage(context.Background(), event.ConversationId, message.ID); err != nil {
				fmt.Println(err)
			}
		} else {
			event.ParentMessageId = message.Meta["threadId"].(string)
			if err := client.ReadThreadMessage(context.Background(), event.ConversationId, event.ParentMessageId, message.ID); err != nil {
				fmt.Println(err)
			}
		}

		if strings.HasPrefix(message.Text, "/") {
			event.UserId = ""
			event.Type = "Command"
			event.Payload, _ = strings.CutPrefix(message.Text, "/")
		} else {
			event.UserId = ""
			event.Type = "Message"
		}

		if client.tokenInfo.UserId != message.Author.ID {
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
