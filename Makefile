.PHONY: build

build-CreateSettings:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler functions/storeSetting.go