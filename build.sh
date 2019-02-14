#!/usr/bin/env bash

VERSION="v1.0.0"
NAME="base64-tool"

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./releases/${NAME}-win32-${VERSION}.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./releases/${NAME}-win64-${VERSION}.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ./releases/${NAME}-darwin32-${VERSION}
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./releases/${NAME}-darwin64-${VERSION}
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./releases/${NAME}-linux32-${VERSION}
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./releases/${NAME}-linux64-${VERSION}