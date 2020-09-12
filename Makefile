.PHONY: list goinit gotidy install run test

.DEFAULT_GOAL := install

SHELL := /bin/bash

PROJECT_NAME = a_star

list:
	# Listing of available make targets
	@(grep -oe '^[[:lower:][:digit:]_\-]\{1,\}:' Makefile | tr -d ':' | uniq)

goinit:
	go mod init github.com/jxtopher/$(PROJECT_NAME)

gotidy:
	go mod tidy

install:
ifneq ($(PACKAGE),)
	go get $(PACKAGE) $(ARGS)
else
	go get ./...
endif

run:
	go run main.go

test: TARGET ?= ...
test:
ifneq ($(TARGET),)
	go test -v ./$(TARGET) $(ARGS)
else
	go test ./$(TARGET) $(ARGS)
endif
