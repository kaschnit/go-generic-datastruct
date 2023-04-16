
SHELL := /bin/bash
.ONESHELL:

GOVERSION := $(shell ./scripts/go-version-utils.sh goModVersion)
GOBIN := $(shell go env GOPATH)/bin
TOOL_FILE := ./tools/tools.go

.PHONY: unittest
unittest: TARGET=./...
unittest:
	go test -coverpkg=$(TARGET) -coverprofile=.profile.cov $(TARGET)
	go tool cover -func .profile.cov

.PHONY: download
download:
	@echo downloading go.mod dependencies
	@go mod download

.PHONY: install-tools
install-tools: download
	@echo installing tools from $(TOOL_FILE)
	@cat $(TOOL_FILE) | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: check
check:
	@echo checking code...
	@$(GOBIN)/golangci-lint run || exit 1
	@echo ok

.PHONY: fix
fix:
	@echo fixing code...
	@$(GOBIN)/golangci-lint run --fix || exit 1
	@echo ok
