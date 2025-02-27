.DEFAULT_GOAL := build

fmt:
	go fmt client.go
.PHONY:fmt

lint: fmt
	golint client.go
.PHONY:lint

vet: fmt
	go vet client.go
.PHONY:vet

build: vet
	go build client.go
.PHONY:build