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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deadline-team/dtalks-bot-api/model"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"github.com/deadline-team/dtalks-bot-api/util"
	"io"
	"net/http"
	"time"
)

const userBasePath = "/api/authorization/users"

var userSrv UserService

type UserService interface {
	// GetUserById
	// Метод для получения пользователей по ID
	GetUserById(ctx context.Context, userId string, fields string) (*userModel.User, error)

	// GetUserAvatarById
	// Метод для получения аватаров пользователей по ID
	GetUserAvatarById(ctx context.Context, userId string, dimension string) ([]byte, error)

	// GetUserAll
	// Метод для получения всех пользователей с фильтрацией
	GetUserAll(ctx context.Context, page model.Pageable, filter userModel.UserFilter, fields string) (*model.Page[userModel.User], error)
}

type userService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewUserService(botBaseParam model.BotBaseParam) UserService {
	if userSrv != nil {
		return userSrv
	}
	userSrv = &userService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return userSrv
}

func (service *userService) GetUserById(ctx context.Context, userId string, fields string) (*userModel.User, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", userBasePath, userId), nil)
	if err != nil {
		return nil, err
	}
	appendUserQueryParams(request, model.Pageable{}, userModel.UserFilter{}, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var user *userModel.User
	if err := json.NewDecoder(response.Body).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) GetUserAvatarById(ctx context.Context, userId string, dimension string) ([]byte, error) {
	if dimension == "" {
		dimension = "128"
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s/avatar?dimension=%s", userBasePath, userId, dimension), nil)
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

func (service *userService) GetUserAll(ctx context.Context, page model.Pageable, filter userModel.UserFilter, fields string) (*model.Page[userModel.User], error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, userBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendUserQueryParams(request, page, filter, fields)

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var users model.Page[userModel.User]
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		return nil, err
	}
	return &users, nil
}

func appendUserQueryParams(request *http.Request, page model.Pageable, filter userModel.UserFilter, fields string) {
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
	if filter.Username != "" {
		request.Form.Set("username", filter.Username)
	}
	if filter.FirstName != "" {
		request.Form.Set("firstName", filter.FirstName)
	}
	if filter.LastName != "" {
		request.Form.Set("lastName", filter.LastName)
	}
	if filter.Email != "" {
		request.Form.Set("email", filter.Email)
	}
	if filter.Search != "" {
		request.Form.Set("search", filter.Search)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
