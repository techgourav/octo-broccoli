#!/bin/bash

# Build protocol
cd protocol
./build.sh
cd ..

# Build src code
cd server
go build -ldflags "-s -w"
cd ../..