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
	calendarEventModel "github.com/deadline-team/dtalks-bot-api/model/calendar_event"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	taskModel "github.com/deadline-team/dtalks-bot-api/model/task"
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

	//TODO GetUserById(ctx context.Context, userId string, fields string) authorizationModel.User
	//TODO GetUserAvatarById(ctx context.Context, userId string, dimension string) []byte
	//TODO GetUserAll(ctx context.Context, page model.Pageable, fields string) model.Page[authorizationModel.User]
	//TODO GetUserAllAdmins(ctx context.Context) []authorizationModel.User
	//TODO FindUserByUsername(ctx context.Context, username string, fields string) authorizationModel.User
	//TODO FindUserByEmail(ctx context.Context, email string, fields string) authorizationModel.User

	//TODO BlockUserById(ctx context.Context, userId string)
	//TODO UnblockUserById(ctx context.Context, userId string)
	//TODO RefreshUserTokenById(ctx context.Context, userId string)
	//TODO DropUserCacheById(ctx context.Context, userId string)
	//TODO LogoutUserById(ctx context.Context, userId string)

	/*---------------------- Conversations ----------------------*/

	//TODO GetConversationById(ctx context.Context, conversationId string, fields string) conversationModel.Conversation
	//TODO GetConversationAvatarById(ctx context.Context, conversationId string, dimension string) []byte

	// GetConversationAll
	// Метод для получения всех открытых диалогов с ботом
	GetConversationAll(ctx context.Context) ([]conversationModel.Conversation, error)

	//TODO CreateConversation(ctx context.Context, dType conversationModel.ConversationDType, conversation conversationModel.Conversation) conversationModel.Conversation
	//TODO UpdateConversation(ctx context.Context, conversation conversationModel.Conversation) conversationModel.Conversation
	//TODO DeleteConversationById(ctx context.Context, conversationId string)
	//TODO AddConversationMemberById(ctx context.Context, conversationId string, userId string)
	//TODO RemoveConversationMemberById(ctx context.Context, conversationId string, userId string)

	/*---------------------- Messages ----------------------*/

	// CreateMessage
	// Метод для создания и отправки сообщения в диалог
	CreateMessage(ctx context.Context, conversationId string, message conversationModel.Message) (*conversationModel.Message, error)

	//TODO GetMessageById(ctx context.Context, messageId string, fields string) conversationModel.Message
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

	// GetLabelById
	// Метод для получения меток по ID
	GetLabelById(ctx context.Context, labelId string, fields string) (*conversationModel.Label, error)

	// GetLabelAll
	// Метод для получения всех меток с фильтрацией
	GetLabelAll(ctx context.Context, page model.Pageable, filter conversationModel.LabelFilter, fields string) (*model.Page[conversationModel.Label], error)

	// CreateLabel
	// Метод для создания меток
	CreateLabel(ctx context.Context, label conversationModel.Label) (*conversationModel.Label, error)

	// UpdateLabel
	// Метод для обновления меток
	UpdateLabel(ctx context.Context, label conversationModel.Label) (*conversationModel.Label, error)

	// DeleteLabelById
	// Метод для удаления меток по ID
	DeleteLabelById(ctx context.Context, labelId string) error

	/*---------------------- Reactions ----------------------*/

	// GetReactionById
	// Метод для получения реакций по ID
	GetReactionById(ctx context.Context, reactionId string, fields string) (*conversationModel.Reaction, error)

	// GetReactionAll
	// Метод для получения всех реакций с фильтрацией
	GetReactionAll(ctx context.Context, page model.Pageable, filter conversationModel.ReactionFilter, fields string) (*model.Page[conversationModel.Reaction], error)

	// CreateReaction
	// Метод для создания реакций
	CreateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error)

	// UpdateReaction
	// Метод для обновления реакций
	UpdateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error)

	// DeleteReactionById
	// Метод для удаления реакций по ID
	DeleteReactionById(ctx context.Context, reactionId string) error

	/*---------------------- Meetings ----------------------*/

	// GetCalendarEventById
	// Метод для получения событий календаря по ID
	GetCalendarEventById(ctx context.Context, meetingId string, fields string) (*calendarEventModel.CalendarEvent, error)

	// GetCalendarEventAll
	// Метод для получения всех событий календаря с фильтрацией
	GetCalendarEventAll(ctx context.Context, filter calendarEventModel.CalendarEventFilter, fields string) ([]calendarEventModel.CalendarEvent, error)

	// CreateCalendarEvent
	// Метод для создания событий календаря
	CreateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error)

	// UpdateCalendarEvent
	// Метод для обновления событий календаря
	UpdateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error)

	// DeleteCalendarEventById
	// Метод для удаления событий календаря по ID
	DeleteCalendarEventById(ctx context.Context, meetingId string) error

	/*---------------------- Tasks ----------------------*/

	// GetTaskById
	// Метод для получения задач по ID
	GetTaskById(ctx context.Context, taskId string, fields string) (*taskModel.Task, error)

	// GetTaskAll
	// Метод для получения всех задач с фильтрацией
	GetTaskAll(ctx context.Context, filter taskModel.TaskFilter, fields string) ([]taskModel.Task, error)

	// CreateTask
	// Метод для создания задач
	CreateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error)

	// UpdateTask
	// Метод для обновления задач
	UpdateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error)

	// DeleteTaskById
	// Метод для удаления задач по ID
	DeleteTaskById(ctx context.Context, taskId string) error

	// ResolveTaskById
	// Метод для маркировки задач как исполненной
	ResolveTaskById(ctx context.Context, taskId string) error
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
