.DEFAULT_GOAL := build

BIN_FILE=go-simple-docker
IMAGE_VERSION = v0.1.1
BUILDER_IMAGE_NAME=gosimpledocker 

build:
	@go build -o "${BIN_FILE}"

build-dockerfile:				## Create dockerfile which acts as build environment
	@docker build --pull -t $(BUILDER_IMAGE_NAME) .

push: ## Push the images to local and remote registry
	@docker push $(IMAGE_NAME):$(IMAGE_VERSION)

clean:
	go clean
	rm --force "cp.out"
	rm --force nohup.out

test:
	go test

check:
	go test

cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out

run:
	./"${BIN_FILE}"

lint:
	golangci-lint run --enable-all

	