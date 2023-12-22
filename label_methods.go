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
	"github.com/deadline-team/dtalks-bot-api/model"
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	"net/http"
)

const labelBasePath = "/api/conversation/labels"

func (client *botAPI) GetLabelById(ctx context.Context, labelId string, fields string) (*conversationModel.Label, error) {
	request, err := client.createRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", labelBasePath, labelId), nil)
	if err != nil {
		return nil, err
	}
	appendLabelQueryParams(request, model.Pageable{}, conversationModel.LabelFilter{}, fields)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var label *conversationModel.Label
	if err := json.NewDecoder(response.Body).Decode(label); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return label, nil
}

func (client *botAPI) GetLabelAll(ctx context.Context, page model.Pageable, filter conversationModel.LabelFilter, fields string) (*model.Page[conversationModel.Label], error) {
	request, err := client.createRequest(ctx, http.MethodGet, labelBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendLabelQueryParams(request, page, filter, fields)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var labels model.Page[conversationModel.Label]
	if err := json.NewDecoder(response.Body).Decode(&labels); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return &labels, nil
}

func (client *botAPI) CreateLabel(ctx context.Context, label conversationModel.Label) (*conversationModel.Label, error) {
	data, err := json.Marshal(&label)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPost, labelBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&label); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &label, err
}

func (client *botAPI) UpdateLabel(ctx context.Context, label conversationModel.Label) (*conversationModel.Label, error) {
	data, err := json.Marshal(&label)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPut, labelBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&label); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &label, err
}

func (client *botAPI) DeleteLabelById(ctx context.Context, labelId string) error {
	request, err := client.createRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", labelBasePath, labelId), nil)
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

func appendLabelQueryParams(request *http.Request, page model.Pageable, filter conversationModel.LabelFilter, fields string) {
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
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
