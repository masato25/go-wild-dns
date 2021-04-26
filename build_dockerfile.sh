#!/bin/sh

GOOS=linux go build
echo "generated binary"
docker build -t go-wild-dns:v0.1 .
