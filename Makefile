VERSION=$(shell git describe --tags --always)

.PHONY: master
# 将远程master覆盖到本地
master:
	git fetch --all
	git reset --hard origin/master
	git pull

.PHONY: init
# init 安装所需包
init:
	go mod tidy
	go get -u github.com/google/wire/cmd/wire@latest


.PHONY: wire
# generate wire code
wire:
	cd cmd/app/ && wire

.PHONY: build
# build application
build:
	rm -rf bin
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: run
# run application
run:
	make wire && go run ./...


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help