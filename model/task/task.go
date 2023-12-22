package task

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
	conversationModel "github.com/deadline-team/dtalks-bot-api/model/conversation"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"time"
)

type Task struct {
	ID           string                          `json:"id,omitempty"`
	Number       string                          `json:"number,omitempty"`
	Description  string                          `json:"description,omitempty"`
	CreateDate   *time.Time                      `json:"createDate,omitempty"`
	Creator      *userModel.User                 `json:"creator,omitempty"`
	Responsible  *userModel.User                 `json:"responsible,omitempty"`
	Deadline     *time.Time                      `json:"deadline,omitempty"`
	Conversation *conversationModel.Conversation `json:"conversation,omitempty"`
	Message      *conversationModel.Message      `json:"message,omitempty"`
	Resolved     bool                            `json:"resolved,omitempty"`
	ResolvedDate *time.Time                      `json:"resolvedDate,omitempty"`
}

type TaskFilter struct {
	IDs            []string `form:"ids"`
	ConversationId string   `form:"conversationId"`
	Resolved       bool     `form:"resolved"`
	Search         string   `form:"search"`
}
