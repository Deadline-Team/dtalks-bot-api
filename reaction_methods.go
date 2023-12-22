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

const reactionBasePath = "/api/conversation/reactions"

func (client *botAPI) GetReactionById(ctx context.Context, reactionId string, fields string) (*conversationModel.Reaction, error) {
	request, err := client.createRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", reactionBasePath, reactionId), nil)
	if err != nil {
		return nil, err
	}
	appendReactionQueryParams(request, model.Pageable{}, conversationModel.ReactionFilter{}, fields)

	response, err := httpClient.Do(request)
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
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return reaction, nil
}

func (client *botAPI) GetReactionAll(ctx context.Context, page model.Pageable, filter conversationModel.ReactionFilter, fields string) (*model.Page[conversationModel.Reaction], error) {
	request, err := client.createRequest(ctx, http.MethodGet, reactionBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendReactionQueryParams(request, page, filter, fields)

	response, err := httpClient.Do(request)
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
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return &reactions, nil
}

func (client *botAPI) CreateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error) {
	data, err := json.Marshal(&reaction)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPost, reactionBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&reaction); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &reaction, err
}

func (client *botAPI) UpdateReaction(ctx context.Context, reaction conversationModel.Reaction) (*conversationModel.Reaction, error) {
	data, err := json.Marshal(&reaction)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPut, reactionBasePath, bytes.NewReader(data))
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
	if err := json.NewDecoder(response.Body).Decode(&reaction); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &reaction, err
}

func (client *botAPI) DeleteReactionById(ctx context.Context, reactionId string) error {
	request, err := client.createRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", reactionBasePath, reactionId), nil)
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
