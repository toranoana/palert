#!/bin/sh
go build -o ./bin/mac/palert palert.go
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/palert palert.go
