reimport:
	go mod tidy

generate:
	test -f $(GOPATH)/bin/easyjson || go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	cd model && $(GOPATH)/bin/easyjson -all avatar.go
	cd model && $(GOPATH)/bin/easyjson -all event.go
	cd model && $(GOPATH)/bin/easyjson -all pageable.go
	cd model && $(GOPATH)/bin/easyjson -all sort.go
	cd model && $(GOPATH)/bin/easyjson -all token.go
	cd model/attachment && $(GOPATH)/bin/easyjson -all attachment.go
	cd model/calendar_event && $(GOPATH)/bin/easyjson -all calendar_event.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all conversation.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all label.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all link.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message_button.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all message_reaction.go
	cd model/conversation && $(GOPATH)/bin/easyjson -all reaction.go
	cd model/scheduler && $(GOPATH)/bin/easyjson -all schduled_task.go
	cd model/user && $(GOPATH)/bin/easyjson -all user.go
	cd model/task && $(GOPATH)/bin/easyjson -all task.go

.PHONY: reimport generate
