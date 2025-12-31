PACKAGE = github.com/wonjinsin/simple-mcp
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BIN = $(BASE_PATH)/bin
BINARY_NAME = bin/mark3labs
MAIN = $(BASE_PATH)/cmd/mark3labs/main.go
GOLINT = $(BIN)/golint
GOBIN = $(shell go env GOPATH)/bin
PKG_LIST = $(shell cd $(BASE_PATH) && cat pkg.list)


ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif

.PHONY: tool
tool:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: build
build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

.PHONY: vet
vet:
	go vet

.PHONY: fmt
fmt:
	go fmt

.PHONY: lint
lint:
	$(GOBIN)/golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: test-all
test-all: test vet fmt lint

.PHONY: init
init: 
	go mod init $(PACKAGE)

.PHONY: tidy
tidy: 
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

# Infrastructure commands
.PHONY: infra-up
infra-up:
	docker compose up -d

.PHONY: infra-down
infra-down:
	docker compose down

.PHONY: start
start: build 
	@$(BINARY_NAME)

.PHONY: all
all: tool init tidy vendor build

.PHONY: clean
clean:; $(info cleaning…) @ 
	@rm -rf vendor mock bin
	@rm -rf go.mod go.sum pkg.list
