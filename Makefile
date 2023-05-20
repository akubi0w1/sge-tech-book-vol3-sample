CURRENT_PATH = $(shell pwd)

# toolchains
TOOLCHAIN_DIR = $(CURRENT_PATH)/.toolchain
TOOLCHAIN_BIN = $(TOOLCHAIN_DIR)/bin

# go
GO_VERSION = 1.19.2
GOROOT = $(TOOLCHAIN_DIR)/go
GOPATH = $(TOOLCHAIN_DIR)/go
GO = $(GOROOT)/bin/go

# protoc
PROTO_PATH = $(CURRENT_PATH)/api/protos
PROTOC_GEN_GO_VERSION = 1.26.0
PROTOC_GEN_GO_GRPC_VERSION = 1.1.0
PROTOC_GEN_DOC = 1.5.1
PROTOC_GEN_GRPC_GATEWAY_VERSION = 2.14.0
PROTOC_GEN_OPENAPI_VERSION = 2.14.0

# buf
BUF_VERSION = 1.15.1
BUF = $(TOOLCHAIN_BIN)/buf

UNAME_S = $(shell uname -s)
ARCH = $(shell uname -m)
ifeq ($(UNAME_S),Linux)
	BUF_PACKAGE = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_S)-$(ARCH)
	GO_PACKAGE = https://go.dev/dl/go$(GO_VERSION).linux-$(ARCH).tar.gz
endif
ifeq ($(UNAME_S),Darwin)
	ifeq ($(ARCH),x86_64) # intel
		BUF_PACKAGE = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_S)-$(ARCH)
		GO_PACKAGE = https://go.dev/dl/go$(GO_VERSION).darwin-amd64.tar.gz	
	else # m1
		BUF_PACKAGE = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_S)-$(ARCH)
		GO_PACKAGE = https://go.dev/dl/go$(GO_VERSION).darwin-arm64.tar.gz
	endif
endif

# #########################################
# buf for local

# usage: buf-help
buf-help: $(BUF)
	$(BUF) --help

# usage: buf-init
buf-init: $(BUF)
	cd $(PROTO_PATH) && $(BUF) mod init

# bufの依存関係などを更新した際に必要
# usage: buf-mod-update
buf-mod-update: $(BUF)
	cd $(PROTO_PATH) && $(BUF) mod update

# usage: buf-build
buf-build: $(BUF)
	$(BUF) build -o tools/protodesc.json $(PROTO_PATH)

# usage: buf-lint
buf-lint: $(BUF)
	$(BUF) lint $(PROTO_PATH)

# usage: buf-generate
buf-generate: $(BUF)
	$(BUF) generate --template $(PROTO_PATH)/buf.gen.yaml $(PROTO_PATH)

# usage: buf-format
buf-format: $(BUF)
	$(BUF) format -w

# install buf
$(BUF): $(GOPATH)/bin/protoc-gen-go $(GOPATH)/bin/protoc-gen-go-grpc $(GOPATH)/bin/protoc-gen-doc $(GOPATH)/bin/protoc-gen-grpc-gateway $(GOPATH)/bin/protoc-gen-openapiv2 $(GOPATH)/bin/protoc-gen-validate
	mkdir -p $(TOOLCHAIN_BIN)
	curl -sSL $(BUF_PACKAGE) -o "$(TOOLCHAIN_BIN)/buf"
	chmod +x "$(TOOLCHAIN_BIN)/buf"

# install protoc-gen-go
$(GOPATH)/bin/protoc-gen-go: $(GOROOT)
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@v$(PROTOC_GEN_GO_VERSION)

## install protoc-gen-go-grpc
$(GOPATH)/bin/protoc-gen-go-grpc: $(GOROOT)
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$(PROTOC_GEN_GO_GRPC_VERSION)

## install protoc-gen-doc
$(GOPATH)/bin/protoc-gen-doc: $(GOROOT)
	$(GO) install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v$(PROTOC_GEN_DOC)

## install protoc-gen-grpc-gateway
$(GOPATH)/bin/protoc-gen-grpc-gateway: $(GOROOT)
	$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v$(PROTOC_GEN_GRPC_GATEWAY_VERSION)

## install protoc-gen-openapiv2
$(GOPATH)/bin/protoc-gen-openapiv2: $(GOROOT)
	$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v$(PROTOC_GEN_OPENAPI_VERSION)

## install protoc-gen-validate
$(GOPATH)/bin/protoc-gen-validate: $(GOROOT)
	$(GO) get -d github.com/envoyproxy/protoc-gen-validate

# ###################################
# go tools

# usage: go-version
go-version: $(GOROOT)
	$(GO) version

# install go
$(GOROOT):
	mkdir -p $(TOOLCHAIN_DIR)
	curl -SL $(GO_PACKAGE) | tar xzC $(TOOLCHAIN_DIR)

go-dbgen: $(GOPATH) $(GOPATH)/bin/sqlboiler $(GOPATH)/bin/sqlboiler-mysql
	cd tools/dbentitygen && $(GO) generate ./...

$(GOPATH)/bin/sqlboiler: $(GOROOT) $(GOPATH)
	$(GO) get github.com/volatiletech/sqlboiler/v4@v4.12.0
	$(GO) get github.com/volatiletech/sqlboiler/v4/boilingcore@v4.12.0
	$(GO) install github.com/volatiletech/sqlboiler/v4

$(GOPATH)/bin/sqlboiler-mysql: $(GOROOT) $(GOPATH)
	$(GO) install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql

go-masterhandlergen:
	- mkdir tools/masterhandlergen/tmp/
	$(BUF) build -o tools/masterhandlergen/tmp/protodesc.json $(PROTO_PATH)
	$(GO) run tools/masterhandlergen/main.go

# ###################################
# cli
cli-subcmdgen:
	- mkdir tools/clisubcmdgen/tmp/
	$(BUF) build -o tools/clisubcmdgen/tmp/protodesc.json $(PROTO_PATH)
	$(GO) run tools/clisubcmdgen/main.go

cli-build:
	$(GO) build -o bin/cli ./cmd/cli/main.go

# ###################################
# docker compose

up:
	docker-compose up --build $(SERVICE) $(OPT)

down:
	docker-compose down --rmi local --volumes --rmi local

# ###################################
# database

db-migrate:
	MYSQL_MASTER_ADDR=localhost:3306 \
	MYSQL_MASTER_PROTOCOL=tcp \
	MYSQL_MASTER_USER=root \
	MYSQL_MASTER_PASSWORD=root \
	MYSQL_MASTER_DB=master \
	MYSQL_SHARD_ADDR=localhost:3306 \
	MYSQL_SHARD_PROTOCOL=tcp \
	MYSQL_SHARD_USER=root \
	MYSQL_SHARD_PASSWORD=root \
	MYSQL_SHARD_DB=shard \
	$(GO) run tools/migrate/main.go
