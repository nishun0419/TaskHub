# GoFlow

Go（Gin）とNext.js 14を使用した、JWT認証・Google認証対応のシンプルな会員管理アプリケーションです。  
Dockerを用いた開発環境構築も対応しています。

---

## 📚 使用技術

### バックエンド
- Go 1.24
- Gin
- GORM
- Goose（マイグレーションツール）
- MySQL

### フロントエンド
- Next.js 14 (App Router)
- Tailwind CSS
- NextAuth.js（Google OAuth認証）
- JWT認証（アクセストークン）

### その他
- Docker / Docker Compose
### ディレクトリ構造
```
.
├── backend/
│   ├── cmd/                # エントリーポイント
│   ├── controllers/        # ハンドラー層
│   ├── models/             # データモデル
│   ├── service/            # ビジネスロジック
│   ├── validators/         # バリデーション
│   └── db/                 # DB接続・マイグレーション
└── frontend/
    ├── app/                # App Router構成
    ├── constants/          # 定数管理
    └── public/             # 静的ファイル
```

---

## 🚀 機能一覧

- ユーザー登録（メールアドレス・パスワード）
- JWT発行によるログイン認証
- GoogleアカウントによるOAuthログイン
- マイページ閲覧（認証後のみ）
- ログアウト処理
- （今後追加予定）リフレッシュトークン対応

---

## 🛠️ 環境構築方法

1. リポジトリをクローン
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
docker compose up
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
docker compose exec frontend sh

# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev
```

## テスト

### バックエンド
```bash
# コンテナ内で実行
docker compose exec backend sh
go test ./...
```

### フロントエンド
```bash
# コンテナ内で実行
docker compose exec frontend sh
npm test
```