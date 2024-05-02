include .env
label ?= local
tag ?= 941355383184.dkr.ecr.us-east-2.amazonaws.com/sca-worker:$(label)

.PHONY: lint
lint:
	go fmt ./...
	golangci-lint run

.PHONY: build
build:
	docker build -t $(tag) --build-arg GIT_TOKEN=$(GITHUB_TOKEN) .

.PHONY: test
test:
	go test ./... -count=1