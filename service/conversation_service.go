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
	"io"
	"net/http"
	"time"
)

const conversationBasePath = "/api/conversation/conversations"

var conversationSrv ConversationService

type ConversationService interface {
	// GetConversationById
	// Метод для получения диалогов по ID
	GetConversationById(ctx context.Context, conversationId string, fields string) (*conversationModel.Conversation, error)

	// GetConversationAvatarById
	// Метод для получения аватаров диалогов по ID
	GetConversationAvatarById(ctx context.Context, conversationId string, dimension string) ([]byte, error)

	// GetConversationAll
	// Метод для получения всех открытых диалогов с ботом
	GetConversationAll(ctx context.Context, page model.Pageable, filter conversationModel.ConversationFilter, fields string) (*model.Page[conversationModel.Conversation], error)

	// CreateChatConversation
	// Метод для создания чата с пользователем
	CreateChatConversation(ctx context.Context, otherMemberId string) (*conversationModel.Conversation, error)

	// CreateGroupConversation
	// Метод для создания группы
	CreateGroupConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error)

	// CreateChannelConversation
	// Метод для создания канала
	CreateChannelConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error)

	// UpdateConversation
	// Метод для обновления диалогов
	UpdateConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error)

	// DeleteConversationById
	// Метод для удаления диалога по ID
	DeleteConversationById(ctx context.Context, conversationId string) error

	// AddConversationMemberById
	// Метод для добавления участника в диалог по userID
	AddConversationMemberById(ctx context.Context, conversationId string, userId string) error

	// RemoveConversationMemberById
	// Метод для удаления участника из диалога по userID
	RemoveConversationMemberById(ctx context.Context, conversationId string, userId string) error
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

func (service *conversationService) GetConversationById(ctx context.Context, conversationId string, fields string) (*conversationModel.Conversation, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", conversationBasePath, conversationId), nil)
	if err != nil {
		return nil, err
	}
	appendConversationQueryParams(request, model.Pageable{}, conversationModel.ConversationFilter{}, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var conversation conversationModel.Conversation
	if err := json.NewDecoder(response.Body).Decode(&conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (service *conversationService) GetConversationAvatarById(ctx context.Context, conversationId string, dimension string) ([]byte, error) {
	if dimension == "" {
		dimension = "128"
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s/avatar?dimension=%s", conversationBasePath, conversationId, dimension), nil)
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	return io.ReadAll(response.Body)
}

func (service *conversationService) GetConversationAll(ctx context.Context, page model.Pageable, filter conversationModel.ConversationFilter, fields string) (*model.Page[conversationModel.Conversation], error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, conversationBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendConversationQueryParams(request, page, filter, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
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
	return &conversationPage, nil
}

func (service *conversationService) CreateChatConversation(ctx context.Context, otherMemberId string) (*conversationModel.Conversation, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, fmt.Sprintf("%s/chat?otherMemberId=%s", conversationBasePath, otherMemberId), nil)
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var conversation conversationModel.Conversation
	if err := json.NewDecoder(response.Body).Decode(&conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (service *conversationService) CreateGroupConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error) {
	data, err := json.Marshal(&conversation)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, fmt.Sprintf("%s/group", conversationBasePath), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (service *conversationService) CreateChannelConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error) {
	data, err := json.Marshal(&conversation)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, fmt.Sprintf("%s/channel", conversationBasePath), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (service *conversationService) UpdateConversation(ctx context.Context, conversation conversationModel.Conversation) (*conversationModel.Conversation, error) {
	data, err := json.Marshal(&conversation)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, conversationBasePath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (service *conversationService) DeleteConversationById(ctx context.Context, conversationId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s", conversationBasePath, conversationId), nil)
	if err != nil {
		return err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

func (service *conversationService) AddConversationMemberById(ctx context.Context, conversationId string, userId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/members/%s", conversationBasePath, conversationId, userId), nil)
	if err != nil {
		return err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

func (service *conversationService) RemoveConversationMemberById(ctx context.Context, conversationId string, userId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/members/%s", conversationBasePath, conversationId, userId), nil)
	if err != nil {
		return err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

func appendConversationQueryParams(request *http.Request, page model.Pageable, filter conversationModel.ConversationFilter, fields string) {
	if page.Page != 0 {
		request.Form.Set("page", fmt.Sprintf("%d", page.Page))
	}
	if page.Size != 0 {
		request.Form.Set("size", fmt.Sprintf("%d", page.Size))
	}
	if page.Sort != nil {
		request.Form.Set("sort", fmt.Sprintf("%s,%s", page.Sort.Field, page.Sort.Order))
	}
	if filter.IDs != nil && len(filter.IDs) > 0 {
		for _, id := range filter.IDs {
			request.Form.Add("ids", id)
		}
	}
	if filter.Name != "" {
		request.Form.Set("name", filter.Name)
	}
	if filter.Visibility != "" {
		request.Form.Set("visibility", string(filter.Visibility))
	}
	if filter.Search != "" {
		request.Form.Set("search", filter.Search)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
