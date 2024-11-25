.PHONY: build test test-integration generate install-paramgen install-tools golangci-lint-install

VERSION=$(shell git describe --tags --dirty --always)

build:
	go build -ldflags "-X 'github.com/conduitio-labs/conduit-connector-zeromq.version=${VERSION}'" -o conduit-connector-zeromq cmd/connector/main.go

test:
	go test $(GOTEST_FLAGS) -race ./...

test-integration:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose.yml up -d
	go test $(GOTEST_FLAGS) -v -race ./...; ret=$$?; \
		docker compose -f test/docker-compose.yml down; \
		exit $$ret

generate:
	go generate ./...

install-paramgen:
	go install github.com/conduitio/conduit-connector-sdk/cmd/paramgen@latest

install-tools:
	@echo Installing tools from tools.go
	@go list -e -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy

lint:
	golangci-lint run -v
