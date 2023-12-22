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
	"context"
	"encoding/json"
	"errors"
	"github.com/deadline-team/dtalks-bot-api/model"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/util"
	"net/http"
	"time"
)

const conversationBasePath = "/api/conversation/conversations"

var conversationSrv ConversationService

type ConversationService interface {
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
}

type conversationService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewConversationService(botBaseParam model.BotBaseParam) ConversationService {
	if conversationSrv != nil {
		return conversationSrv
	}
	conversationSrv = &conversationService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return conversationSrv
}

func (service *conversationService) GetConversationAll(ctx context.Context) ([]conversationModel.Conversation, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, conversationBasePath, nil)
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
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
