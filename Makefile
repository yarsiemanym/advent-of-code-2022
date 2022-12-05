.DEFAULT_GOAL := build
SOURCE := $(wildcard *.go **/*.go)
GO_PATH := $(shell go env GOPATH)
INSTALL_PATH := $(GO_PATH)/bin/advent-of-code-2022
LOG_LEVEL ?= warn

.PHONY:
session:
ifndef AOC_SESSION_TOKEN
	$(error AOC_SESSION_TOKEN is undefined)
endif

.PHONY:
test: session
	go test common/*.go
	@#go test day00/*.go
	go test day01/*.go
	go test day02/*.go
	go test day03/*.go
	go test day04/*.go
	go test day05/*.go
	go test day06/*.go

.PHONY:
build: advent-of-code-2022

advent-of-code-2022: $(SOURCE)
	go build

.PHONY:
clean: 
	go clean

.PHONY:
run: build session
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 $(DAY) $(INPUT_FILE)

.PHONY:
run-all: build session
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 0
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 1
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 2
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 3
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 4
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 5
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2022 6

.PHONY:
install: $(INSTALL_PATH)

$(INSTALL_PATH): $(SOURCE)
	go install 

.PHONY:
uninstall:
	rm -f $(INSTALL_PATH)

.PHONY:
setup: go.mod
	go mod tidy
	go mod download

.PHONY:
go.mod: /usr/local/go/bin/go ~/.go
	go mod init

/usr/local/go/bin/go:
	sudo wget -c https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

~/.go:
	echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee $(HOME)/.go
	echo -e "\n. \$$HOME/.go" | tee -a $(HOME)/.bashrc
	