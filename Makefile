
NAME = iserv2go

SHELL = /bin/sh
RM = rm
GO = go

OUT_DIR = build

VERSION = $(shell git describe --tags --abbrev=0)
COMMIT = $(shell git rev-parse --short $(VERSION))

GOFLAGS = -ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT_SHA=$(COMMIT)"

# build

build:
	go build -v $(GOFLAGS) .

# linux
build-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-amd64 .

build-i386:
	GOOS=linux GOARCH=386 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-386 .

build-arm:
	GOOS=linux GOARCH=arm $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-arm .

build-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-linux-arm64 .

# windows
build-win-amd64:
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-windows-amd64.exe

build-win-i386:
	GOOS=windows GOARCH=386 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-windows-386.exe

# mac
build-mac:
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUT_DIR)/$(NAME)-darwin-arm64

build-all: build-amd64 build-i386 build-arm build-arm64 build-win-amd64 build-win-i386 build-mac

# install

install:
	$(GO) install .

# clean

clean:
	$(RM) -rf build

# info

info:
	@echo $(VERSION) $(COMMIT)
