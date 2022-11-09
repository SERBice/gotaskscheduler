#!/bin/bash

mkdir -p bin
GOOS=windows GOARCH=amd64 go build -o bin/windows-app-64.exe test.go
GOOS=windows GOARCH=386 go build -o bin/windows-app-32.exe test.go
GOOS=windows GOARCH=arm64 go build -o bin/windows-app-arm64.exe test.go

GOOS=linux GOARCH=amd64 go build -o bin/linux-app-64 test.go
GOOS=linux GOARCH=386 go build -o bin/linux-app-32 test.go
GOOS=linux GOARCH=arm64 go build -o bin/linux-app-arm64 test.go

GOOS=darwin GOARCH=amd64 go build -o bin/mac-app-64 test.go
GOOS=darwin GOARCH=386 go build -o bin/mac-app-32 test.go
GOOS=darwin GOARCH=arm64 go build -o bin/mac-app-arm64 test.go