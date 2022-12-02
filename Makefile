
NAME = iserv2go

SHELL = /bin/sh
RM = rm
GO = go

OUT_DIR = build

AMD64 = amd64
I386 = 386
ARM = arm

build:
	go build .

build-amd64:
	GOOS=linux GOARCH=$(AMD64) $(GO) build -o $(OUT_DIR)/$(NAME)-linux-$(AMD64) .

build-i386:
	GOOS=linux GOARCH=$(I386) $(GO) build -o $(OUT_DIR)/$(NAME)-linux-$(I386) .

build-arm:
	GOOS=linux GOARCH=$(ARM) $(GO) build -o $(OUT_DIR)/$(NAME)-linux-$(ARM) .

build-win:
	GOOS=windows GOARCH=$(AMD64) $(GO) build -o $(OUT_DIR)/$(NAME)-windows-$(AMD64).exe .

build-all: build-amd64 build-i386 build-arm build-win

install:
	$(GO) install .

clean:
	$(RM) -rf build
