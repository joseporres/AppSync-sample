.PHONY: build

build-PlusTenFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/PlusTen.go

build-GetInvitationTemplateFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/get-invitation-template.go/main.go

build-SendInvitationFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/send-invitation.go/main.go

	