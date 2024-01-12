package service

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
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"github.com/deadline-team/dtalks-bot-api/util"
	"net/http"
	"time"
)

const reactionBasePath = "/api/conversation/reactions"

var reactionSrv ReactionService

type ReactionService interface {
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
}

type reactionService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewReactionService(botBaseParam model.BotBaseParam) ReactionService {
	if reactionSrv != nil {
		return reactionSrv
	}
	reactionSrv = &reactionService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return reactionSrv
}

func (service *reactionService) GetReactionById(ctx context.Context, reactionId string, fields string) (*conversationModel.Reaction, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", reactionBasePath, reactionId), nil)
	if err != nil {
		return nil, err
	}
	appendReactionQueryParams(request, model.Pageable{}, conversationModel.ReactionFilter{}, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var reaction *conversationModel.Reaction
	if err := json.NewDecoder(response.Body).Decode(reaction); err != nil {
		return nil, err
	}
	return reaction, nil
}

func (service *reactionService) GetReactionAll(ctx context.Context, page model.Pageable, filter conversationModel.ReactionFilter, fields string) (*model.Page[conversationModel.Reaction], error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, reactionBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendReactionQueryParams(request, page, filter, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var reactions model.Page[conversationModel.Reaction]
	if err := json.NewDecoder(response.Body).Decode(&reactions); err != nil {
		return nil, err
	}
	return &reactions, nil
}

func (service *reactionService) CreateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error) {
	data, err := json.Marshal(&reaction)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, reactionBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&reaction); err != nil {
		return nil, err
	}

	return &reaction, err
}

func (service *reactionService) UpdateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error) {
	data, err := json.Marshal(&reaction)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, reactionBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&reaction); err != nil {
		return nil, err
	}

	return &reaction, err
}

func (service *reactionService) DeleteReactionById(ctx context.Context, reactionId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s", reactionBasePath, reactionId), nil)
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

func appendReactionQueryParams(request *http.Request, page model.Pageable, filter conversationModel.ReactionFilter, fields string) {
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
	if filter.Value != "" {
		request.Form.Set("value", filter.Value)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
