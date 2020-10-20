#!/bin/sh

env GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o bin/sammy.exe cmd/sammy.go

echo ""
echo "Built bin/sammy.exe"
