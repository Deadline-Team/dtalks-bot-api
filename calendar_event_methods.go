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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	calendarEventModel "github.com/deadline-team/dtalks-bot-api/model/calendar_event"
	"net/http"
	"time"
)

const calendarEventBasePath = "/api/calendarEvent/calendarEvents"

func (client *botAPI) GetCalendarEventById(ctx context.Context, calendarEventId string, fields string) (*calendarEventModel.CalendarEvent, error) {
	request, err := client.createRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", calendarEventBasePath, calendarEventId), nil)
	if err != nil {
		return nil, err
	}
	appendCalendarEventQueryParams(request, calendarEventModel.CalendarEventFilter{}, fields)

	response, err := httpClient.Do(request)
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

func (client *botAPI) GetCalendarEventAll(ctx context.Context, filter calendarEventModel.CalendarEventFilter, fields string) ([]calendarEventModel.CalendarEvent, error) {
	request, err := client.createRequest(ctx, http.MethodGet, calendarEventBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendCalendarEventQueryParams(request, filter, fields)

	response, err := httpClient.Do(request)
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

func (client *botAPI) CreateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error) {
	data, err := json.Marshal(&calendarEvent)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPost, calendarEventBasePath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
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

func (client *botAPI) UpdateCalendarEvent(ctx context.Context, calendarEvent calendarEventModel.CalendarEvent) (*calendarEventModel.CalendarEvent, error) {
	data, err := json.Marshal(&calendarEvent)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPut, calendarEventBasePath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
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

func (client *botAPI) DeleteCalendarEventById(ctx context.Context, calendarEventId string) error {
	request, err := client.createRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", calendarEventBasePath, calendarEventId), nil)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
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
