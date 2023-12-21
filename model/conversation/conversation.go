package conversation

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
	"github.com/deadline-team/dtalks-bot-api/model"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"time"
)

type ConversationDType string

const (
	ConversationType ConversationDType = "Conversation"
	ChatType         ConversationDType = "Chat"
	GroupType        ConversationDType = "Group"
	ChannelType      ConversationDType = "Channel"
)

type Visibility string

const (
	Public  Visibility = "PUBLIC"
	Private Visibility = "PRIVATE"
)

type Conversation struct {
	ID           string              `json:"id,omitempty"`
	Type         []ConversationDType `json:"type,omitempty"`
	Name         string              `json:"name,omitempty"`
	Avatar       *model.Avatar       `json:"avatar,omitempty"`
	Visibility   Visibility          `json:"visibility,omitempty"`
	CreateDate   *time.Time          `json:"createDate,omitempty"`
	Owner        *userModel.User     `json:"owner,omitempty"`
	Admins       []*userModel.User   `json:"admins,omitempty"`
	Members      []*userModel.User   `json:"members,omitempty"`
	MutedMembers []*userModel.User   `json:"mutedMembers,omitempty"`
	Messages     []*Message          `json:"messages,omitempty"`
	Pinned       []*Message          `json:"pinned,omitempty"`
	UnreadCount  int                 `json:"unreadCount,omitempty"`
	MessageCount int                 `json:"messageCount,omitempty"`
	UnreadThread []*Message          `json:"unreadThread,omitempty"`
}
