#!/usr/bin/env bash


wire() {
    echo "wiring dependencies..."
    go run -mod=mod github.com/google/wire/cmd/wire gen github.com/gnanasuriyan/go-message-server/internal
}

initDatabase() {
    echo "initializing database..."
    docker exec -i mysql_db bash -c "mysql -uroot -proot -e 'CREATE DATABASE IF NOT EXISTS message_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;'"
    goose -dir ./migrations mysql "root:root@tcp(localhost:3306)/message_db?parseTime=true" up  # Apply the migrations
}

if [ "$1" == "wire" ]; then
    wire
elif [ "$1" == "init-db" ]; then
  initDatabase
fi
