#!/bin/bash

go test -count=1 -coverprofile=cover.out ./...
go tool cover -func=cover.out
rm cover.out
