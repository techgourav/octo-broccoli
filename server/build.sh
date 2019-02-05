#!/bin/bash

# Build protocol
cd protocol
./build.sh
cd ..

# Build src code
cd src
go build -ldflags "-s -w"
cd ..