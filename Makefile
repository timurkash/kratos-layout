GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
KRATOS_THIRD_PARTY_DIR=$(GOPATH)/src/github.com/go-kratos/kratos/third_party

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

.PHONY: errors
# generate errors code
errors:
	@protoc \
		-I=. \
		--proto_path=$(KRATOS_THIRD_PARTY_DIR) \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		$(API_PROTO_FILES)

tidy:
	@go mod tidy -compat=1.17

.PHONY: config
# generate internal proto
config:
	@protoc \
		-I=. \
		-I=$(KRATOS_THIRD_PARTY_DIR) \
		--go_out=paths=source_relative:. \
		$(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	@protoc \
		-I=. \
		-I=$(KRATOS_THIRD_PARTY_DIR) \
 		--go_out=paths=source_relative:. \
 		--go-grpc_out=paths=source_relative:. \
		$(API_PROTO_FILES)

.PHONY: dc
dc:
	@docker-compose up -d

.PHONY: build
build: config wire
	@mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: run
run: dc build
	./bin/$(shell ls bin)

.PHONY: wire
wire:
	@wire gen ./cmd/server

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all: api errors config generate

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: test
test:
	@GO111MODULE=on go test -v -race -cover ./internal/service
	@GO111MODULE=on go test ./internal/service -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@firefox coverage.html
