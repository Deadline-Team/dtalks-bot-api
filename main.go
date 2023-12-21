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
	attachmentModel "github.com/deadline-team/dtalks-bot-api/model/attachment"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/util"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type BotAPI interface {
	// GetEventChannel
	// Метод для получения канала в который складываются все события с сервера
	GetEventChannel() chan model.Event

	/*---------------------- Attachments ----------------------*/

	// CreateAttachment
	// Метод для создания вложения на сервере
	CreateAttachment(ctx context.Context, fileName string, data []byte) (*attachmentModel.Attachment, error)
	//TODO GetAttachmentById(ctx context.Context, attachmentId string) ([]byte, error)
	//TODO GetAttachmentMetaById(ctx context.Context, attachmentId string) (*attachmentModel.Attachment, error)

	/*---------------------- Users ----------------------*/

	//TODO GetUserById(ctx context.Context, id string, fields string) authorizationModel.User
	//TODO GetUserAvatarById(ctx context.Context, id string, dimension string) []byte
	//TODO GetUserAll(ctx context.Context, page model.Pageable, fields string) model.Page[authorizationModel.User]
	//TODO GetUserAllAdmins(ctx context.Context) []authorizationModel.User
	//TODO FindUserByUsername(ctx context.Context, username string, fields string) authorizationModel.User
	//TODO FindUserByEmail(ctx context.Context, email string, fields string) authorizationModel.User

	//TODO BlockUserById(ctx context.Context, id string)
	//TODO UnblockUserById(ctx context.Context, id string)
	//TODO RefreshUserTokenById(ctx context.Context, id string)
	//TODO DropUserCacheById(ctx context.Context, id string)
	//TODO LogoutUserById(ctx context.Context, id string)

	/*---------------------- Conversations ----------------------*/

	// GetConversationAll
	// Метод для получения всех открытых диалогов с ботом
	GetConversationAll(ctx context.Context) ([]conversationModel.Conversation, error)

	//TODO GetConversationById(ctx context.Context, id string, fields string) conversationModel.Conversation
	//TODO GetConversationAvatarById(ctx context.Context, id string, dimension string) []byte
	//TODO CreateConversation(ctx context.Context, dType conversationModel.ConversationDType, conversation conversationModel.Conversation) conversationModel.Conversation
	//TODO UpdateConversation(ctx context.Context, conversation conversationModel.Conversation) conversationModel.Conversation
	//TODO DeleteConversationById(ctx context.Context, conversationId string)
	//TODO AddConversationMemberById(ctx context.Context, conversationId string, userId string)
	//TODO RemoveConversationMemberById(ctx context.Context, conversationId string, userId string)

	/*---------------------- Messages ----------------------*/

	// CreateMessage
	// Метод для создания и отправки сообщения в диалог
	CreateMessage(ctx context.Context, conversationId string, message conversationModel.Message) (*conversationModel.Message, error)

	//TODO GetMessageById(ctx context.Context, id string, fields string) conversationModel.Message
	//TODO GetMessageAll(ctx context.Context, conversationId string) model.Page[conversationModel.Message]
	//TODO UpdateMessage(ctx context.Context, conversationId string, messageId string, text string) conversationModel.Message
	//TODO DeleteMessageById(ctx context.Context, conversationId string, messageId string)
	//TODO PinMessage(ctx context.Context, conversationId string, messageId string)
	//TODO UnpinMessage(ctx context.Context, conversationId string, messageId string)
	//TODO ReadMessage(ctx context.Context, conversationId string, messageId string)
	//TODO AddLabelToMessage(ctx context.Context, conversationId string, messageId string, labelId string)
	//TODO RemoveLabelToMessage(ctx context.Context, conversationId string, messageId string, labelId string)
	//TODO ReactMessage(ctx context.Context, conversationId string, messageId string, reactionId string)

	//TODO GetThreadMessageAll(ctx context.Context, conversationId string, parentMessageId string) conversationModel.Message
	//TODO CreateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, message conversationModel.Message) conversationModel.Message
	//TODO UpdateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, text string) conversationModel.Message
	//TODO DeleteThreadMessageById(ctx context.Context, conversationId string, parentMessageId string, messageId string)
	//TODO ReadThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string)
	//TODO AddLabelToThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string)
	//TODO RemoveLabelToThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string)
	//TODO ReactThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, reactionId string)

	/*---------------------- Labels ----------------------*/

	//TODO GetLabelById(ctx context.Context, id string, fields string) conversationModel.Label
	//TODO GetLabelAll(ctx context.Context, page model.Pageable, fields string) model.Page[conversationModel.Label]
	//TODO CreateLabel(ctx context.Context, label conversationModel.Label) conversationModel.Label
	//TODO UpdateLabel(ctx context.Context, label conversationModel.Label) conversationModel.Label
	//TODO DeleteLabelById(ctx context.Context, id string)
	//TODO FindLabelByName(ctx context.Context, name string) conversationModel.Label

	/*---------------------- Reactions ----------------------*/

	//TODO GetById(ctx context.Context, id string, fields string) conversationModel.Reaction
	//TODO GetAll(ctx context.Context, page model.Pageable, fields string) model.Page[conversationModel.Reaction]
	//TODO Create(ctx context.Context, reaction conversationModel.Reaction) conversationModel.Reaction
	//TODO Update(ctx context.Context, reaction conversationModel.Reaction) conversationModel.Reaction
	//TODO DeleteById(ctx context.Context, id string)
	//TODO FindByValue(ctx context.Context, value string) conversationModel.Reaction

	/*---------------------- Meetings ----------------------*/

	//TODO GetMeetingById(ctx context.Context, id string, fields string) calendarEventModel.CalendarEvent
	//TODO GetMeetingAll(ctx context.Context, fields string) []calendarEventModel.CalendarEvent
	//TODO CreateMeeting(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) calendarEventModel.CalendarEvent
	//TODO UpdateMeeting(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) calendarEventModel.CalendarEvent
	//TODO DeleteMeetingById(ctx context.Context, id string)

	/*---------------------- Tasks ----------------------*/

	//TODO GetTaskById(ctx context.Context, id string, fields string) taskModel.Task
	//TODO GetTaskAll(ctx context.Context, filter taskModel.TaskFilter, fields string) []taskModel.Task
	//TODO CreateTask(ctx context.Context, task taskModel.Task) taskModel.Task
	//TODO UpdateTask(ctx context.Context, task taskModel.Task) taskModel.Task
	//TODO DeleteTaskById(ctx context.Context, id string)
	//TODO ResolveTaskById(ctx context.Context, id string)
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
