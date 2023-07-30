# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

OS ?= $(shell uname)
ARCH ?= $(shell uname -m)

GOOS ?= $(shell echo "$(OS)" | tr '[:upper:]' '[:lower:]')
GOARCH_x86_64 = amd64
GOARCH_aarch64 = arm64
GOARCH_arm64 = arm64
GOARCH ?= $(shell echo "$(GOARCH_$(ARCH))")

REVISION := dev.$(shell echo $(CIRCLE_SHA1) | head -c 8 | git rev-parse --short=8 HEAD)

OUTPUT_DIR := ./
OUTPUT_BIN := lab-api-teams

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint against code.
	golangci-lint run ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: test
test: ## Run tests.
	go test ./... -coverprofile cover.out

.PHONY: static 
static: lint fmt test

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT_DIR)/$(OUTPUT_BIN)  cmd/lab-api-teams/main.go
	chmod +x $(OUTPUT_DIR)/$(OUTPUT_BIN)
	docker build -t docker.io/twdps/lab-api-teams:$(REVISION) .

.PHONY: push
push:
	docker push docker.io/twdps/lab-api-teams:$(REVISION)

.PHONY: increment
increment:
	-./semver bump patch $(shell git tag --points-at  HEAD~1) | xargs -I{} git tag v{}




