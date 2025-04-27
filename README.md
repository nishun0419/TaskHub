# GoFlow

GoFlowは、Go言語とNext.jsを使用したフルスタックのWebアプリケーションです。

## システム構成

### バックエンド (Go)
- フレームワーク: Gin
- データベース: MySQL
- 認証: JWT
- マイグレーション: Goose

### フロントエンド (Next.js)
- フレームワーク: Next.js 14
- UI: Tailwind CSS
- 認証: NextAuth.js (Google認証),JWT

## ディレクトリ構成

```
.
├── backend/                 # バックエンド（Go）
│   ├── cmd/                # メインロジック
│   ├── controllers/        # HTTPハンドラー
│   ├── models/            # データモデル
│   ├── service/           # ビジネスロジック
│   ├── validators/        # バリデーション
│   └── db/                # データベース関連
└── frontend/              # フロントエンド（Next.js）
    ├── app/              # アプリケーションコード
    ├── constants/       # 定数ファイル
    └── public/          # 静的ファイル
```

## 環境要件

- Docker
- Docker Compose
- Go 1.24以上
- Node.js 20以上

## セットアップ手順

1. リポジトリのクローン
```bash
git clone https://github.com/nishun0419/goflow.git
cd goflow
cp frontend/env frontend/.env.local
```
※googleログインをする場合は.env.localに必要な情報を入れてください

2. Dockerコンテナのビルド
```bash
docker compose build 
```
3. アプリケーションの起動
```bash
# 開発環境
docker-compose up
```

4. アプリケーションにアクセス
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

## 開発環境での作業

### バックエンド
```bash
# コンテナに入る
docker compose exec backend sh

# マイグレーションの実行
GOOSE_DBSTRING="${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" go run github.com/pressly/goose/v3/cmd/goose@latest -dir db/migrations up

# アプリケーションの起動
go run cmd/main.go
```

### フロントエンド
```bash
# コンテナに入る
docker-compose exec frontend sh

# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev
```

## テスト

### バックエンド
```bash
# コンテナ内で実行
docker-compose exec backend sh
go test ./...
```

### フロントエンド
```bash
# コンテナ内で実行
docker-compose exec frontend sh
npm test
```

## デプロイ

1. 本番環境用の環境変数を設定
```bash
cp backend/.env.prod backend/.env
cp frontend/.env.prod frontend/.env
```

2. 本番環境用のビルドと起動
```bash
ENV=prod docker-compose up --build
```