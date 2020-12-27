BIN_NAME = sbc

VERSION_TAG = $(shell git describe --tags --abbrev=0 --always)
VERSION_COMMIT = $(shell git rev-parse --short HEAD)
VERSION_DATE = $(shell git show -s --format=%cI HEAD)
VERSION = -X main.versionTag=$(VERSION_TAG) -X main.versionCommit=$(VERSION_COMMIT) -X main.versionDate=$(VERSION_DATE)

.PHONY: help
help:
	@ cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: mod
mod: ## Get dependency packages
	go mod tidy

.PHONY: build
build: ## Build binary file
	go build -a -ldflags "$(VERSION)" -o ./bin/${BIN_NAME} ./cmd/...

.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: race
race: ## Run data race detector
	go test -race ./...

.PHONY: cover
cover: ## Generate code coverage statistics
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html;

.PHONY: lint
lint: ## Lint the files
	golangci-lint version
	golangci-lint run
