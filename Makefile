# go option
GO             ?= go
GOIMPORTS      ?= goimports
PKG            := $(GO) mod vendor
APP_VERSION    := 1.0.0-$(shell git rev-parse HEAD)
LDFLAGS        := -w -s
GOFLAGS        :=
GOPRIVATE      := 
GOPROXY        := https://proxy.golang.org
GO_EXTRA_FLAGS := -v -tags=or_dev
TAGS           :=
BINDIR         := $(CURDIR)/bin
BIN_NAME       := grocerylistsbot
PKGDIR         := github.com/jaredallard/$(BIN_NAME)
CGO_ENABLED    := 1
TOOL_DEPS      := ${GO} ${GOIMPORTS}

.PHONY: default
default: build

.PHONY: pre-commit
pre-commit: fmt

.PHONY: generate-schema
generate-schema:
	go run github.com/facebookincubator/ent/cmd/entc generate ./ent/schema

.PHONY: build
build: gogenerate gobuild

.PHONY: test
test:
	GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) ./scripts/test.sh

.PHONY: docs
docs:
	@echo "Not done yet"

.PHONY: dep
dep:
	@echo " ===> Installing dependencies via <=== "
	GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) $(PKG)

.PHONY: gogenerate
gogenerate:
	GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) $(GO) generate ./...

.PHONY: gobuild
gobuild:
	@echo " ===> building releases in ./bin/... <=== "
	GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) CGO_ENABLED=$(CGO_ENABLED) $(GO) build -o $(BINDIR)/$(BIN_NAME) -ldflags "$(LDFLAGS)" $(GO_EXTRA_FLAGS) $(PKGDIR)/cmd/$(BIN_NAME)

.PHONY: docker-build
docker-build:
	@echo " ===> building docker image <==="
	@ssh-add -L
	@echo " ===> If you run into credential issues, ensure that your key is in your SSH agent (ssh-add <ssh-key-path>) <==="
	DOCKER_BUILDKIT=1 docker build --ssh default -t gcr.io/outreach-docker/authz -f deployments/authz/Dockerfile . --build-arg VERSION=${APP_VERSION}

.PHONY: fmt
fmt:
	@echo " ===> Running goimports <==="
	find  . -path ./vendor -prune -o -type f -name '*.go' -print | xargs -n 1 ${GOIMPORTS} -w
