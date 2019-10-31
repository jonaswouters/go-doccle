SHELL := /usr/bin/env bash

export PROJECT = go-doccle

all: retriever-build

retriever-run:
	go run ./cmd/retriever

retriever-build:
	go build ./cmd/retriever

test:
	go test ./...

run: retriever-run