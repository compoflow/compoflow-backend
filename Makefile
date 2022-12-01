# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: build

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

build: fmt vet ## Build manager binary.
	go build -o bin/backend main.go

run: fmt vet ## Run code from your host.
	go run ./main.go

IMG ?= tjcadworkflow/backend
docker-build:
	docker buildx build -t ${IMG} . --load

docker-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMG} . --push