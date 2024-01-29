package conversation

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
	"github.com/deadline-team/dtalks-bot-api/model"
	attachmentModel "github.com/deadline-team/dtalks-bot-api/model/attachment"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"time"
)

type MessageSubType string

const (
	Base      MessageSubType = "base"
	Date      MessageSubType = "date"
	Technical MessageSubType = "technical"
	Call      MessageSubType = "call"
	Mail      MessageSubType = "mail"
	Poll      MessageSubType = "poll"
)

type Message struct {
	ID                string                        `json:"id,omitempty"`
	SubType           MessageSubType                `json:"subType,omitempty"`
	CreateDate        *time.Time                    `json:"createDate,omitempty"`
	Text              string                        `json:"text,omitempty"`
	Author            *userModel.User               `json:"author,omitempty"`
	Reply             *Message                      `json:"reply,omitempty"`
	Forward           *Message                      `json:"forward,omitempty"`
	Thread            []*Message                    `json:"thread,omitempty"`
	ThreadCount       int                           `json:"threadCount,omitempty"`
	ThreadUnreadCount int                           `json:"threadUnreadCount,omitempty"`
	Meta              model.Meta                    `json:"meta,omitempty"`
	Edited            bool                          `json:"edited,omitempty"`
	EditDate          *time.Time                    `json:"editDate,omitempty"`
	UnreadUsers       []*userModel.User             `json:"unread,omitempty"`
	Labels            []*Label                      `json:"labels,omitempty"`
	MessageReactions  []*MessageReaction            `json:"messageReactions,omitempty"`
	Attachments       []*attachmentModel.Attachment `json:"attachments,omitempty"`
	Links             []*Link                       `json:"links,omitempty"`
	Read              bool                          `json:"read,omitempty"`
	ReadDate          *time.Time                    `json:"readDate,omitempty"`
	CallConferenceId  string                        `json:"callConferenceId,omitempty"`
	History           model.Meta                    `json:"history,omitempty"`
	Buttons           MessageButtons                `json:"buttons,omitempty"`
	Deleted           bool                          `json:"deleted,omitempty"`
	DeletedDate       *time.Time                    `json:"deletedDate,omitempty"`
}
