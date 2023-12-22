package service

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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deadline-team/dtalks-bot-api/model"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/util"
	"net/http"
	"time"
)

var messageSrv MessageService

type MessageService interface {
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
}

type messageService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewMessageService(botBaseParam model.BotBaseParam) MessageService {
	if messageSrv != nil {
		return messageSrv
	}
	messageSrv = &messageService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return messageSrv
}

func (service *messageService) CreateMessage(ctx context.Context, conversationId string, message conversationModel.Message) (*conversationModel.Message, error) {
	data, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, fmt.Sprintf("%s/%s/messages", conversationBasePath, conversationId), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &message, err
}
