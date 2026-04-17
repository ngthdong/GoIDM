VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
PROJECT_NAME := GoIDM

all: generate build-all

.PHONY: generate 
generate:
	protoc -I=. \
		--go_out=internal/generated \
		--go-grpc_out=internal/generated \
		--grpc-gateway_out=internal/generated \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out . \
		--openapiv2_opt generate_unbound_methods=true \
		--validate_out="lang=go:internal/generated" \
		api/go_load.proto

	wire ./internal/wiring	

.PHONY: build
build:
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME) \
		cmd/$(PROJECT_NAME)/*.go

.PHONY: clean
clean:
	rm -rf build/

.PHONY: docker-compose-dev-up
docker-compose-dev-up:
	docker-compose -f deployments/docker-compose.dev.yml up -d

.PHONY: docker-compose-dev-down
docker-compose-dev-down:
	docker-compose -f deployments/docker-compose.dev.yml down

.PHONY: run-server
run-server:
	go run cmd/*.go server

.PHONY: lint
lint:
	golangci-lint run ./... 