GO ?= go
CMD ?= help

.PHONY: vendor_install install format watch build release_build run

default: build

watch: 
	@air $(CMD)

run:
	@$(GO) run main.go $(CMD)

build:
	@$(GO) build -o wo .

release_build:
	@echo [Wo Release] - Building for darwin
	@GOOS=darwin GOARCH=amd64 $(GO) build -o bin/darwin/wo_amd64
	@GOOS=darwin GOARCH=arm64 $(GO) build -o bin/darwin/wo_arm64
	@echo [Wo Release] - Building for BSDs
	@GOOS=freebsd GOARCH=386 $(GO) build -o bin/freebsd/wo_386
	@GOOS=freebsd GOARCH=amd64 $(GO) build -o bin/freebsd/wo_amd64
	@GOOS=freebsd GOARCH=arm $(GO) build -o bin/freebsd/wo_arm
	@GOOS=freebsd GOARCH=arm64 $(GO) build -o bin/freebsd/wo_arm64
	@echo [Wo Release] - Building for linux
	@GOOS=linux GOARCH=386 $(GO) build -o bin/linux/wo_386
	@GOOS=linux GOARCH=amd64 $(GO) build -o bin/linux/wo_amd64
	@GOOS=linux GOARCH=arm $(GO) build -o bin/linux/wo_arm
	@GOOS=linux GOARCH=arm64 $(GO) build -o bin/linux/wo_arm64
	@echo [Wo Release] - Building for windows
	@GOOS=windows GOARCH=386 $(GO) build -o bin/win/wo_386.exe
	@GOOS=windows GOARCH=amd64 $(GO) build -o bin/win/wo_amd64.exe
	@GOOS=windows GOARCH=arm $(GO) build -o bin/win/wo_arm.exe

package: release_build
	@./build/builder.sh

vendor_install:
	@$(GO) mod vendor 

install:
	@$(GO) mod download

format:
	@$(GO) fmt ./**/**/*.go
