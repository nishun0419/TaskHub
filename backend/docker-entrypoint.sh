#!/bin/sh
set -e

# マイグレーションを実行
echo "Running database migrations..."
GOOSE_DBSTRING="${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" go run github.com/pressly/goose/v3/cmd/goose@latest -dir db/migrations up

# アプリケーションを起動
echo "Starting application..."
exec go run cmd/main.go