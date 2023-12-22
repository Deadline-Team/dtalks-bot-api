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
	"github.com/deadline-team/dtalks-bot-api/model"
	"net/http"
	"time"
)

const userBasePath = "/api/authorization/users"

var userSrv UserService

type UserService interface {
	//TODO GetUserById(ctx context.Context, userId string, fields string) authorizationModel.User
	//TODO GetUserAvatarById(ctx context.Context, userId string, dimension string) []byte
	//TODO GetUserAll(ctx context.Context, page model.Pageable, fields string) model.Page[authorizationModel.User]
	//TODO GetUserAllAdmins(ctx context.Context) []authorizationModel.User
	//TODO FindUserByUsername(ctx context.Context, username string, fields string) authorizationModel.User
	//TODO FindUserByEmail(ctx context.Context, email string, fields string) authorizationModel.User

	//TODO BlockUserById(ctx context.Context, userId string)
	//TODO UnblockUserById(ctx context.Context, userId string)
	//TODO RefreshUserTokenById(ctx context.Context, userId string)
	//TODO DropUserCacheById(ctx context.Context, userId string)
	//TODO LogoutUserById(ctx context.Context, userId string)
}

type userService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewUserService(botBaseParam model.BotBaseParam) UserService {
	if userSrv != nil {
		return userSrv
	}
	userSrv = &userService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return userSrv
}
