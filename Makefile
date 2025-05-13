ROOT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR	:= $(ROOT_DIR)/bin/
SERVICE_NAME_API ?= terraform-provider-nodeshift

default: install

generate:
	go generate ./...

install:
	go install .

test:
	go test -count=1 -parallel=4 ./...

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...

gen_docs:
	tfplugindocs generate && tfplugindocs validate

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(SERVICE_NAME_API) main.go

.PHONY: docker-build
docker-build:
	$(call in_docker,make build)
	echo "path for local bin is $(BIN_DIR)"

.PHONY: lint
lint:
	golangci-lint run --config configs/.golangci.yml

define in_docker
	docker run --rm \
		-v $(PWD):/app \
		-w /app \
		golang:1.23 $1
endef