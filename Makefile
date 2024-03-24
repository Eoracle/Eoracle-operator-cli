COMMIT_HASH=$(shell git rev-parse HEAD)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			 grep -o '".*"' | sed 's/"//g' 2> /dev/null || echo v0.0.5)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD | tr -d '\040\011\012\015\n')
TIME=$(shell date)

# Set the source directory
SRC_DIR=cmd/operatorcli

# Set the Docker image names
DOCKER_IMAGE_OPR_CLI_AMD64=eoracle/opr_cli:v$(VERSION)-amd64
DOCKER_IMAGE_OPR_CLI_ARM64=eoracle/opr_cli:v$(VERSION)-arm64


.PHONY: check-git
check-git:
	@which git > /dev/null || (echo "git is not installed. Please install and try again."; exit 1)

.PHONY: check-go
check-go:
	@which go > /dev/null || (echo "Go is not installed.. Please install and try again."; exit 1)

.PHONY: build-oprcli
build-oprcli: check-go check-git
	go build -o bin/oprcli -ldflags="\
    	-X 'github.com/Eoracle/core-go/versioning.Version=$(VERSION)' \
		-X 'github.com/Eoracle/core-go/versioning.Commit=$(COMMIT_HASH)'\
		-X 'github.com/Eoracle/core-go/versioning.Branch=$(BRANCH)'\
		-X 'github.com/Eoracle/core-go/versioning.BuildTime=$(TIME)'" \
	cmd/operatorcli/main.go

.PHONY: build-oprcli-arm64
build-oprcli-arm64: check-go check-git
	GOOS=linux GOARCH=arm64 GOHOSTARCH="arm64" go build -o bin/oprcli-arm64 -ldflags="\
		-X 'github.com/Eoracle/core-go/versioning.Version=$(VERSION)' \
		-X 'github.com/Eoracle/core-go/versioning.Commit=$(COMMIT_HASH)'\
		-X 'github.com/Eoracle/core-go/versioning.Branch=$(BRANCH)'\
		-X 'github.com/Eoracle/core-go/versioning.BuildTime=$(TIME)'" \
	cmd/operatorcli/main.go

.PHONY: build-oprcli-amd64
build-oprcli-amd64: check-go check-git
	GOOS=linux GOARCH=amd64 GOHOSTARCH="amd64" go build -o bin/oprcli-amd64 -ldflags="\
		-X 'github.com/Eoracle/core-go/versioning.Version=$(VERSION)' \
		-X 'github.com/Eoracle/core-go/versioning.Commit=$(COMMIT_HASH)'\
		-X 'github.com/Eoracle/core-go/versioning.Branch=$(BRANCH)'\
		-X 'github.com/Eoracle/core-go/versioning.BuildTime=$(TIME)'" \
	cmd/operatorcli/main.go

# Build Docker images for AMD64 and ARM64
.PHONY: docker-build-oprcli
docker-build-oprcli: docker-build-oprcli-amd64 docker-build-oprcli-arm64

.PHONY: docker-build-oprcli-amd64
docker-build-oprcli-amd64:
	docker build -t $(DOCKER_IMAGE_OPR_CLI_AMD64) --platform linux/amd64 . -f Dockerfile-oprcli

.PHONY: docker-build-oprcli-arm64
docker-build-oprcli-arm64:
	docker build -t $(DOCKER_IMAGE_OPR_CLI_ARM64) --platform linux/arm64 . -f Dockerfile-oprcli

# Push Docker images
.PHONY: docker-push-oprcli-no-latest
docker-push-oprcli-no-latest:
	docker push $(DOCKER_IMAGE_OPR_CLI_AMD64)
	docker push $(DOCKER_IMAGE_OPR_CLI_ARM64)

.PHONY: docker-push-oprcli-latest
docker-push-oprcli-latest:
	docker tag $(DOCKER_IMAGE_OPR_CLI_AMD64) eoracle/opr_cli:latest
	docker push eoracle/opr_cli:latest
	# Clean target, removes the binaries

.PHONY: clean
clean:
	rm -rf $(BINARY_DIR)
