reimport:
	go mod tidy

generate:
	test -f $(GOPATH)/bin/easyjson || go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	cd model && $(GOPATH)/bin/easyjson -all event.go
	cd model && $(GOPATH)/bin/easyjson -all pageable.go
	cd model && $(GOPATH)/bin/easyjson -all sort.go
	cd model && $(GOPATH)/bin/easyjson -all token.go
	cd model/attachment && $(GOPATH)/bin/easyjson -all attachment.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all authority.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all avatar.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all base_user_status.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all device.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all role.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all user.go
	cd model/authorization && $(GOPATH)/bin/easyjson -all user_settings.go
	cd model/calendar_event && $(GOPATH)/bin/easyjson -all calendar_event.go
	cd model/conference && $(GOPATH)/bin/easyjson -all room.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all conversation.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all label.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message_button.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message_reaction.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all reaction.go
	cd model/favorite && $(GOPATH)/bin/easyjson -all favorite.go
	cd model/task && $(GOPATH)/bin/easyjson -all task.go

.PHONY: reimport generate
