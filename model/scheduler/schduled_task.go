package model

import (
	"github.com/deadline-team/dtalks-bot-api/model"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"time"
)

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

type ScheduledTaskDType string

const (
	ScheduledTaskType               ScheduledTaskDType = "ScheduledTask"
	DelayedMessageScheduledTaskType ScheduledTaskDType = "DelayedMessageScheduledTask"
)

type ScheduledTask struct {
	ID         string               `json:"id,omitempty"`
	Type       []ScheduledTaskDType `json:"type"`
	CreateDate *time.Time           `json:"createDate,omitempty"`
	Creator    *userModel.User      `json:"creator,omitempty"`
	Deadline   *time.Time           `json:"deadline,omitempty"`
	Object     model.Meta           `json:"object,omitempty"`
	Meta       model.Meta           `json:"meta,omitempty"`
}
type ScheduledTaskFilter struct {
	IDs            []string `form:"ids"`
	ConversationId string   `form:"conversationId"`
	Resolved       bool     `form:"resolved"`
	Search         string   `form:"search"`
	OwnershipID    string
}
