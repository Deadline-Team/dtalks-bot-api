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

const linkBasePath = "/api/conversation/links"

var linkSrv LinkService

type LinkService interface {
	// GetLinkById
	// Метод для получения ссылок по ID
	GetLinkById(ctx context.Context, linkId string, fields string) (*conversationModel.Link, error)

	// GetLinkAll
	// Метод для получения всех ссылок с фильтрацией
	GetLinkAll(ctx context.Context, page model.Pageable, filter conversationModel.LinkFilter, fields string) (*model.Page[conversationModel.Link], error)

	// CreateLink
	// Метод для создания ссылок
	CreateLink(ctx context.Context, link conversationModel.Link) (*conversationModel.Link, error)

	// UpdateLink
	// Метод для обновления ссылок
	UpdateLink(ctx context.Context, link conversationModel.Link) (*conversationModel.Link, error)

	// DeleteLinkById
	// Метод для удаления ссылок по ID
	DeleteLinkById(ctx context.Context, linkId string) error
}

type linkService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewLinkService(botBaseParam model.BotBaseParam) LinkService {
	if linkSrv != nil {
		return linkSrv
	}
	linkSrv = &linkService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return linkSrv
}

func (service *linkService) GetLinkById(ctx context.Context, linkId string, fields string) (*conversationModel.Link, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", linkBasePath, linkId), nil)
	if err != nil {
		return nil, err
	}
	appendLinkQueryParams(request, model.Pageable{}, conversationModel.LinkFilter{}, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var link *conversationModel.Link
	if err := json.NewDecoder(response.Body).Decode(link); err != nil {
		return nil, err
	}
	return link, nil
}

func (service *linkService) GetLinkAll(ctx context.Context, page model.Pageable, filter conversationModel.LinkFilter, fields string) (*model.Page[conversationModel.Link], error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, linkBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendLinkQueryParams(request, page, filter, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var links model.Page[conversationModel.Link]
	if err := json.NewDecoder(response.Body).Decode(&links); err != nil {
		return nil, err
	}
	return &links, nil
}

func (service *linkService) CreateLink(ctx context.Context, link conversationModel.Link) (*conversationModel.Link, error) {
	data, err := json.Marshal(&link)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, linkBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&link); err != nil {
		return nil, err
	}

	return &link, err
}

func (service *linkService) UpdateLink(ctx context.Context, link conversationModel.Link) (*conversationModel.Link, error) {
	data, err := json.Marshal(&link)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, linkBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&link); err != nil {
		return nil, err
	}

	return &link, err
}

func (service *linkService) DeleteLinkById(ctx context.Context, linkId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s", linkBasePath, linkId), nil)
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

func appendLinkQueryParams(request *http.Request, page model.Pageable, filter conversationModel.LinkFilter, fields string) {
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
