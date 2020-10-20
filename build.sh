#!/bin/bash

echo ""

if [ "$1" == "debug" ]; then
    env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Debug=true" -o bin/sammy-debug.exe cmd/sammy.go && echo "Built bin/sammy-debug.exe"
else
    env GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o bin/sammy.exe cmd/sammy.go && echo "Built bin/sammy.exe"
fi

