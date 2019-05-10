PWD := $(shell pwd)

all: build build-win

build:
	@echo "Building convert-csv binary to './convert-csv'"
	@go build ./cmd/convert-csv.go

build-win:
	@echo "Building convert-csv binary to './convert-csv.exe'"
	@env GOOS=windows GOARCH=amd64 go build ./cmd/convert-csv.go