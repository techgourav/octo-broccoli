#!/bin/bash

# For JavaScript and TypeScript
protoc -I=. --js_out=import_style=commonjs:js ./todo.proto

# For Golang
protoc -I=. --go_out=plugins=grpc:go ./todo.proto