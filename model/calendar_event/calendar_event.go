package calendar_event

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
	attachmentModel "github.com/deadline-team/dtalks-bot-api/model/attachment"
	userModel "github.com/deadline-team/dtalks-bot-api/model/user"
	"time"
)

type TransparencyType string

const (
	BusyTransparencyType      TransparencyType = "Busy"
	AvailableTransparencyType TransparencyType = "Available"
)

type VisibilityType string

const (
	PublicVisibilityType  VisibilityType = "Public"
	PrivateVisibilityType VisibilityType = "Private"
)

type MemberStatusType string

const (
	WaitAnswerMemberStatusType MemberStatusType = "WaitAnswer"
	DeclinedMemberStatusType   MemberStatusType = "Declined"
	TentativeMemberStatusType  MemberStatusType = "Tentative"
	AcceptedMemberStatusType   MemberStatusType = "Accepted"
)

type CalendarEvent struct {
	ID              string                        `json:"id,omitempty"`
	CreateDate      *time.Time                    `json:"createDate,omitempty"`
	UpdateDate      *time.Time                    `json:"updateDate,omitempty"`
	Name            string                        `json:"name,omitempty"`
	Description     string                        `json:"description,omitempty"`
	Location        string                        `json:"location,omitempty"`
	Organizer       *userModel.User               `json:"organizer,omitempty"`
	StartDate       *time.Time                    `json:"startDate,omitempty"`
	EndDate         *time.Time                    `json:"endDate,omitempty"`
	Transparency    TransparencyType              `json:"transparency,omitempty"`
	Visibility      VisibilityType                `json:"visibility,omitempty"`
	Members         []*userModel.User             `json:"members,omitempty"`
	ExternalMembers []string                      `json:"externalMembers,omitempty"`
	MembersStatuses MembersStatusesMap            `json:"membersStatuses,omitempty"`
	Attachments     []*attachmentModel.Attachment `json:"attachments,omitempty"`
}

type CalendarEventFilter struct {
	IDs             []string   `form:"ids"`
	AllUsers        bool       `form:"allUsers"`
	PeriodStartDate *time.Time `form:"periodStartDate"`
	PeriodEndDate   *time.Time `form:"periodEndDate"`
	Search          string     `form:"search"`
}
