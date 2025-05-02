#Railway用のDockerfile
# Golang 1.23の公式イメージを使用
FROM golang:1.23-alpine

# 作業ディレクトリを設定
WORKDIR /app

# キャッシュ利用で効率化するために別でコピー
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY backend/ .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# エントリポイントスクリプトをコピーして実行権限を付与
COPY backend/docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# エントリポイントを設定
ENTRYPOINT ["docker-entrypoint.sh"]

# ポートを公開
EXPOSE 8080

