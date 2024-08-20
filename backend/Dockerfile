# Golang 1.23の公式イメージを使用
FROM golang:1.23-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goのモジュールファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# コンテナ起動時に実行するコマンド
CMD ["./main"]

