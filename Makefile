APP_NAME := currency-converter
CMD_DIR := cmd/converter
BUILD_DIR := bin
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

ARGS ?= ""

.PHONY: build
build:
	GOGC=off go build -v -o $(BUILD_DIR)/$(APP_NAME) ./$(CMD_DIR)

.PHONY: mocks
mocks:
	go generate ./...

.PHONY: test
test:
	go test ./... -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
