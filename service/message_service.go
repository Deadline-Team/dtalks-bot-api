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
	// GetMessageAll
	// Метод для получения всех сообщений в диалоге
	GetMessageAll(ctx context.Context, conversationId string) (*model.Page[conversationModel.Message], error)

	// CreateMessage
	// Метод для создания и отправки сообщений в диалог
	CreateMessage(ctx context.Context, conversationId string, message conversationModel.Message) (*conversationModel.Message, error)

	// UpdateMessage
	// Метод для изменения сообщений в диалоге
	UpdateMessage(ctx context.Context, conversationId string, messageId string, text string) (*conversationModel.Message, error)

	// DeleteMessageById
	// Метод для удаления сообщений из диалога по ID
	DeleteMessageById(ctx context.Context, conversationId string, messageId string) error

	// PinMessage
	// Метод для закрепления сообщений в диалоге по ID
	PinMessage(ctx context.Context, conversationId string, messageId string) error

	// UnpinMessage
	// Метод для удаления закрепления сообщения из диалога по ID
	UnpinMessage(ctx context.Context, conversationId string, messageId string) error

	// AddLabelToMessage
	// Метод для добавления метки к сообщениям в диалоге по ID
	AddLabelToMessage(ctx context.Context, conversationId string, messageId string, labelId string) error

	// RemoveLabelFromMessage
	// Метод для удаления метки к сообщениям в диалоге по ID
	RemoveLabelFromMessage(ctx context.Context, conversationId string, messageId string, labelId string) error

	// ReactMessage
	// Метод для добавления/удаления реакции к сообщениям в диалоге по ID
	ReactMessage(ctx context.Context, conversationId string, messageId string, reactionId string) error

	// GetThreadMessageAll
	// Метод для получения всех сообщений в потоке
	GetThreadMessageAll(ctx context.Context, conversationId string, parentMessageId string) ([]conversationModel.Message, error)

	// CreateThreadMessage
	// Метод для создания и отправки сообщений в поток
	CreateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, message conversationModel.Message) (*conversationModel.Message, error)

	// UpdateThreadMessage
	// Метод для изменения сообщений в потоке
	UpdateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, text string) (*conversationModel.Message, error)

	// DeleteThreadMessageById
	// Метод для удаления сообщений из потока по ID
	DeleteThreadMessageById(ctx context.Context, conversationId string, parentMessageId string, messageId string) error

	// AddLabelToThreadMessage
	// Метод для добавления метки к сообщениям в потоке по ID
	AddLabelToThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string) error

	// RemoveLabelFromThreadMessage
	// Метод для удаления метки к сообщениям в потоке по ID
	RemoveLabelFromThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string) error

	// ReactThreadMessage
	// Метод для добавления/удаления реакции к сообщениям в потоке по ID
	ReactThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, reactionId string) error
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

func (service *messageService) GetMessageAll(ctx context.Context, conversationId string) (*model.Page[conversationModel.Message], error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s/messages", conversationBasePath, conversationId), nil)
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

	var messagePage model.Page[conversationModel.Message]
	if err := json.NewDecoder(response.Body).Decode(&messagePage); err != nil {
		return nil, err
	}
	return &messagePage, nil
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
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}

	return &message, err
}

func (service *messageService) UpdateMessage(ctx context.Context, conversationId string, messageId string, text string) (*conversationModel.Message, error) {
	message := conversationModel.Message{Text: text}
	data, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s", conversationBasePath, conversationId, messageId), bytes.NewReader(data))
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

	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}
	return &message, err
}

func (service *messageService) DeleteMessageById(ctx context.Context, conversationId string, messageId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/messages/%s", conversationBasePath, conversationId, messageId), nil)
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

func (service *messageService) PinMessage(ctx context.Context, conversationId string, messageId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/pin", conversationBasePath, conversationId, messageId), nil)
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

func (service *messageService) UnpinMessage(ctx context.Context, conversationId string, messageId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/messages/%s/pin", conversationBasePath, conversationId, messageId), nil)
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

func (service *messageService) AddLabelToMessage(ctx context.Context, conversationId string, messageId string, labelId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/label/%s", conversationBasePath, conversationId, messageId, labelId), nil)
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

func (service *messageService) RemoveLabelFromMessage(ctx context.Context, conversationId string, messageId string, labelId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/messages/%s/label/%s", conversationBasePath, conversationId, messageId, labelId), nil)
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

func (service *messageService) ReactMessage(ctx context.Context, conversationId string, messageId string, reactionId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/reaction/%s", conversationBasePath, conversationId, messageId, reactionId), nil)
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

func (service *messageService) GetThreadMessageAll(ctx context.Context, conversationId string, parentMessageId string) ([]conversationModel.Message, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s/messages/%s/thread", conversationBasePath, conversationId, parentMessageId), nil)
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

	var messages []conversationModel.Message
	if err := json.NewDecoder(response.Body).Decode(&messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (service *messageService) CreateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, message conversationModel.Message) (*conversationModel.Message, error) {
	data, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, fmt.Sprintf("%s/%s/messages/%s/thread", conversationBasePath, conversationId, parentMessageId), bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}

	return &message, err
}

func (service *messageService) UpdateThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, text string) (*conversationModel.Message, error) {
	message := conversationModel.Message{Text: text}
	data, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/thread/%s", conversationBasePath, conversationId, parentMessageId, messageId), bytes.NewReader(data))
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

	if err := json.NewDecoder(response.Body).Decode(&message); err != nil {
		return nil, err
	}
	return &message, err
}

func (service *messageService) DeleteThreadMessageById(ctx context.Context, conversationId string, parentMessageId string, messageId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/messages/%s/thread/%s", conversationBasePath, conversationId, parentMessageId, messageId), nil)
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

func (service *messageService) AddLabelToThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/thread/%s/label/%s", conversationBasePath, conversationId, parentMessageId, messageId, labelId), nil)
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

func (service *messageService) RemoveLabelFromThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, labelId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s/messages/%s/thread/%s/label/%s", conversationBasePath, conversationId, parentMessageId, messageId, labelId), nil)
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

func (service *messageService) ReactThreadMessage(ctx context.Context, conversationId string, parentMessageId string, messageId string, reactionId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, fmt.Sprintf("%s/%s/messages/%s/thread/%s/reaction/%s", conversationBasePath, conversationId, parentMessageId, messageId, reactionId), nil)
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
