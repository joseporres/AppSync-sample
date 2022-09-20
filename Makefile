.PHONY: build

build-StoreSettingsFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/store/storeSetting.go


build-GetInvitationTemplateFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/get-template-email/getInvitationTemplate.go
	
build-SendInvitationFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/send-invitation-email/sendInvitationEmail.go

build-CheckEmailUserTypeFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/check-email-usertype/main.go

build-CreateAppsProcessFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/agregar-data/main.go
	
build-UpdateForceStatus:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/update_force_status/main.go

build-SendTemporalPassword:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/send_temporal_password/main.go



