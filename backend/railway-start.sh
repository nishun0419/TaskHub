#!/bin/sh
set -e

# アプリケーションを起動
echo "Starting application..."
exec go run cmd/main.go