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
	taskModel "github.com/deadline-team/dtalks-bot-api/model/task"
	"github.com/deadline-team/dtalks-bot-api/util"
	"net/http"
	"time"
)

const taskBasePath = "/api/task/tasks"

var taskSrv TaskService

type TaskService interface {
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

type taskService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewTaskService(botBaseParam model.BotBaseParam) TaskService {
	if taskSrv != nil {
		return taskSrv
	}
	taskSrv = &taskService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return taskSrv
}

func (service *taskService) GetTaskById(ctx context.Context, taskId string, fields string) (*taskModel.Task, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", taskBasePath, taskId), nil)
	if err != nil {
		return nil, err
	}
	appendTaskQueryParams(request, taskModel.TaskFilter{}, fields)

	response, err := service.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var task *taskModel.Task
	if err := json.NewDecoder(response.Body).Decode(task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return task, nil
}

func (service *taskService) GetTaskAll(ctx context.Context, filter taskModel.TaskFilter, fields string) ([]taskModel.Task, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, taskBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendTaskQueryParams(request, filter, fields)

	response, err := service.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var tasks []taskModel.Task
	if err := json.NewDecoder(response.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (service *taskService) CreateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error) {
	data, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, taskBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &task, err
}

func (service *taskService) UpdateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error) {
	data, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, taskBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &task, err
}

func (service *taskService) DeleteTaskById(ctx context.Context, taskId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s", taskBasePath, taskId), nil)
	if err != nil {
		return err
	}

	response, err := service.httpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	if err = response.Body.Close(); err != nil {
		return err
	}
	return nil
}

func (service *taskService) ResolveTaskById(ctx context.Context, taskId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/resolve", taskBasePath, taskId), nil)
	if err != nil {
		return err
	}

	response, err := service.httpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	if err = response.Body.Close(); err != nil {
		return err
	}
	return nil
}

func appendTaskQueryParams(request *http.Request, filter taskModel.TaskFilter, fields string) {
	if filter.IDs != nil && len(filter.IDs) > 0 {
		for _, id := range filter.IDs {
			request.Form.Add("ids", id)
		}
	}
	if filter.ConversationId != "" {
		request.Form.Set("conversationId", filter.ConversationId)
	}
	if filter.Resolved {
		request.Form.Set("resolved", fmt.Sprintf("%t", filter.Resolved))
	}
	if filter.Search != "" {
		request.Form.Set("search", filter.Search)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
