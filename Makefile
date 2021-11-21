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
	@./build/binary-builder.sh

package:
	@./build/package-builder.sh

vendor_install:
	@$(GO) mod vendor 

install:
	@$(GO) mod download

format:
	@$(GO) fmt ./**/**/*.go
