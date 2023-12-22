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
	calendarEventModel "github.com/deadline-team/dtalks-bot-api/model/calendar_event"
	"github.com/deadline-team/dtalks-bot-api/util"
	"net/http"
	"time"
)

const calendarEventBasePath = "/api/calendarEvent/calendarEvents"

var calendarEventSrv CalendarEventService

type CalendarEventService interface {
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
}

type calendarEventService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewCalendarEventService(botBaseParam model.BotBaseParam) CalendarEventService {
	if calendarEventSrv != nil {
		return calendarEventSrv
	}
	calendarEventSrv = &calendarEventService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return calendarEventSrv
}

func (service *calendarEventService) GetCalendarEventById(ctx context.Context, calendarEventId string, fields string) (*calendarEventModel.CalendarEvent, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", calendarEventBasePath, calendarEventId), nil)
	if err != nil {
		return nil, err
	}
	appendCalendarEventQueryParams(request, calendarEventModel.CalendarEventFilter{}, fields)

	response, err := service.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var calendarEvent *calendarEventModel.CalendarEvent
	if err := json.NewDecoder(response.Body).Decode(calendarEvent); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return calendarEvent, nil
}

func (service *calendarEventService) GetCalendarEventAll(ctx context.Context, filter calendarEventModel.CalendarEventFilter, fields string) ([]calendarEventModel.CalendarEvent, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, calendarEventBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendCalendarEventQueryParams(request, filter, fields)

	response, err := service.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var calendarEvents []calendarEventModel.CalendarEvent
	if err := json.NewDecoder(response.Body).Decode(&calendarEvents); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return calendarEvents, nil
}

func (service *calendarEventService) CreateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error) {
	data, err := json.Marshal(&calendarEvent)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, calendarEventBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&calendarEvent); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &calendarEvent, err
}

func (service *calendarEventService) UpdateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error) {
	data, err := json.Marshal(&calendarEvent)
	if err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPut, calendarEventBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&calendarEvent); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &calendarEvent, err
}

func (service *calendarEventService) DeleteCalendarEventById(ctx context.Context, calendarEventId string) error {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodDelete, fmt.Sprintf("%s/%s", calendarEventBasePath, calendarEventId), nil)
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

func appendCalendarEventQueryParams(request *http.Request, filter calendarEventModel.CalendarEventFilter, fields string) {
	if filter.IDs != nil && len(filter.IDs) > 0 {
		for _, id := range filter.IDs {
			request.Form.Add("ids", id)
		}
	}
	if filter.AllUsers {
		request.Form.Set("allUsers", fmt.Sprintf("%t", filter.AllUsers))
	}
	if filter.PeriodStartDate != nil {
		request.Form.Set("periodStartDate", filter.PeriodStartDate.Format(time.RFC3339))
	}
	if filter.PeriodEndDate != nil {
		request.Form.Set("periodEndDate", filter.PeriodEndDate.Format(time.RFC3339))
	}
	if filter.Search != "" {
		request.Form.Set("search", filter.Search)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
