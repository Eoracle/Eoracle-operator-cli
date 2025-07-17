# Default commands and configurations
GO ?= go
V ?= 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# Directories and files
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
BASE_DIR := $(patsubst %/,%,$(dir $(MAKEFILE_PATH)))
SRC_DIR = $(BASE_DIR)
VERSION = $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
-                        grep -o '".*"' | sed 's/"//g' 2> /dev/null || echo v0.0.5)

# Docker image
DOCKER_IMAGE = eoracle/opr-cli

# Linter version
GOLANGCI_LINT_VERSION = v1.62.0

.PHONY: help
help: ## Display this help message
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# --- Dependency Checks ---
.PHONY: check-deps
check-deps: check-go check-git check-fmt ## Verify all required dependencies are installed

.PHONY: check-go
check-go: ## Ensure Go is installed
	@which $(GO) > /dev/null || (echo "Go is not installed. Please install it and try again."; exit 1)

.PHONY: check-git
check-git: ## Ensure Git is installed
	@which git > /dev/null || (echo "Git is not installed. Please install it and try again."; exit 1)

.PHONY: check-fmt
check-fmt: ## Ensure formatting tools are installed
	@which gofmt > /dev/null || (echo "gofmt is not installed. Please install it and try again."; exit 1)
	@which goimports > /dev/null || (echo "goimports is not installed. Please install it and try again."; exit 1)

# --- Linting ---
.PHONY: lint
lint: fmt ## Run golangci-lint on all source files
	$(info $(M) Running golangci-lint...)
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint not found, installing..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION); \
	fi
	$Q golangci-lint run

# --- Formatting ---
.PHONY: fmt
fmt: check-fmt ## Format source files using gofmt and goimports
	$(info $(M) Running gofmt and goimports...)
	@sh $(BASE_DIR)/build/goimports.sh

# --- Testing ---
TIMEOUT ?= 15
ARGS ?= # Optional test arguments
PKGS = $(or $(PKG),$(shell cd $(BASE_DIR) && env GO111MODULE=on $(GO) list ./...))
TESTPKGS = $(shell cd $(BASE_DIR) && env GO111MODULE=on $(GO) list -f \
			  '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			  $(PKGS))

TEST_TARGETS := test-verbose test-race
.PHONY: $(TEST_TARGETS)

# General test rule
.PHONY: test
test: lint ## Run all tests
	$(info $(M) Running $(NAME:%=% )tests...)
	$Q $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

# Test variants
test-verbose: 	ARGS=-v 		## Run tests in verbose mode with coverage reporting
test-race: 	  	ARGS=-race 		## Run tests with race detector

$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test

# --- Dependency Management ---
.PHONY: tidy
tidy: ; $(info $(M) Running go get -u ./... and go mod tidy) @ ## Run go get -u ./... and go mod tidy, then sync workspace.
	@($(GO) get -u ./... && $(GO) mod tidy) || exit 1

# --- Build ---
.PHONY: version
version: ## Print the application version
	@echo $(VERSION)

.PHONY: build
build: build-oprcli ## Build the oprcli binary

.PHONY: build-oprcli
build-oprcli: check-go ## Build oprcli binary
	$(info $(M) Building oprcli binary...)
	$Q $(GO) build -o $(BASE_DIR)/bin/oprcli $(SRC_DIR)/main.go

.PHONY: build-oprcli-%
build-oprcli-%: check-go ## Build oprcli binary for a specific platform
	$(info $(M) Building oprcli binary for $*...)
	$Q GOOS=linux GOARCH=$* $(GO) build -o $(BASE_DIR)/bin/oprcli-$* $(SRC_DIR)/main.go

# --- Docker ---
.PHONY: docker-build
docker-build: docker-build-amd64 docker-build-arm64 ## Build Docker images for all platforms

.PHONY: docker-build-%
docker-build-%: ## Build Docker image for a specific platform
	$(info $(M) Building Docker image for $*...)
	$Q DOCKER_BUILDKIT=1 docker build \
		--build-arg="BUILDPLATFORM=linux/$*" \
		-t $(DOCKER_IMAGE):$(VERSION)-$* --platform linux/$* . -f $(BASE_DIR)/build/Dockerfile

# Push Docker images
.PHONY: docker-push
docker-push:
	docker push $(DOCKER_IMAGE):$(VERSION)-amd64
	docker push $(DOCKER_IMAGE):$(VERSION)-arm64
	docker manifest create $(DOCKER_IMAGE):$(VERSION) --amend $(DOCKER_IMAGE):$(VERSION)-amd64 --amend $(DOCKER_IMAGE):$(VERSION)-arm64
	docker manifest push $(DOCKER_IMAGE):$(VERSION)

.PHONY: docker-push-latest
docker-push-latest:
	docker tag $(DOCKER_IMAGE):$(VERSION)-amd64 $(DOCKER_IMAGE):latest-amd64
	docker tag $(DOCKER_IMAGE):$(VERSION)-arm64 $(DOCKER_IMAGE):latest-arm64
	docker push $(DOCKER_IMAGE):latest-amd64
	docker push $(DOCKER_IMAGE):latest-arm64
	docker manifest create $(DOCKER_IMAGE):latest --amend $(DOCKER_IMAGE):latest-amd64 --amend $(DOCKER_IMAGE):latest-arm64
	docker manifest push $(DOCKER_IMAGE):latest
	# Clean target, removes the binaries

.PHONY: clean
clean:
	rm -rf $(BINARY_DIR)

generate-bindings:
	@echo "Generating bindings"
	@go generate ./...
