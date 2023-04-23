
SHELL := /bin/bash
.ONESHELL:

GOMINVERSION := $(shell ./scripts/go-version-utils.sh goModVersion)
GOVERSION := $(shell which go > /dev/null && go env GOVERSION || echo 0)

ifeq (lt, $(shell ./scripts/go-version-utils.sh compare $(GOVERSION) $(GOMINVERSION)))
$(error "Must have go >= $(GOMINVERSION) (found $(GOVERSION))")
endif

GOBIN := $(shell go env GOPATH)/bin
TOOL_FILE := ./tools/tools.go
COVERAGE_PROFILE := .profile.cov

.PHONY: unittest
unittest: TARGET=./...
unittest: VERBOSE=false
unittest:
	@go test -coverpkg=$(TARGET) -coverprofile=$(COVERAGE_PROFILE) $(TARGET) -v=$(VERBOSE)
	@go tool cover -func=$(COVERAGE_PROFILE)

.PHONY: coverage-html
coverage-html:
	@go tool cover -html=$(COVERAGE_PROFILE)

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
