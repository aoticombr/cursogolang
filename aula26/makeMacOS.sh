#!/bin/bash
export GOOS=darwin
export GOARCH=amd64

go build .
unset GOOS