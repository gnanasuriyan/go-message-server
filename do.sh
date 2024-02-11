#!/usr/bin/env bash


wire() {
    echo "wiring dependencies..."
    go run -mod=mod github.com/google/wire/cmd/wire gen github.com/gnanasuriyan/go-message-server/internal
}

if [ "$1" == "wire" ]; then
    wire
fi

