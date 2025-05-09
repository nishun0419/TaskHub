# チーム管理付きTODOアプリ TaskHub

---

## 1. プロジェクト概要

Go【Gin】をバックエンド、Next.js【14】をフロントエンドとし、
**チーム管理機能付きTODOリストサービス**を開発する。

ユーザーはログイン後、チームを作成し、
チームメンバーとTODOを共有・管理できる。

デプロイ済みで画面の確認も可能です。

https://task-hub-wheat.vercel.app/login

構成図
![goflow構成図.jpg](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/146391/cdf47a98-ab3b-444e-b75c-1ddb86e5740f.jpeg)

チーム作成
![Image](https://github.com/user-attachments/assets/0d2322ec-df2d-4f2f-8f8b-b651e294d18f)
TODO追加
![Image](https://github.com/user-attachments/assets/5cfffa56-9be2-4349-9277-1231dcc1206e)

## 2. 対象ユーザー

- 複数人でタスク管理をしたいユーザー
- チーム単位でプロジェクトを運用したいユーザー


## 3. 機能要件

### 3.1 認証機能

- メールアドレス・パスワードによるログイン/新規登録
- JWTを利用したトークン認証
- ログイン状態の維持

### 3.2 チーム機能

- チーム作成
- チーム一覧表示
- チーム詳細閲覧
- チームメンバー管理 (オプション)

### 3.3 TODO管理機能

- TODO新規作成
- TODO一覧表示
- TODO編集
- TODO削除
- TODO完了チェック
- 自分のTODOのみ編集/削除可能にする

### 3.4 招待機能 (オプション)

- チームに新しいメンバーを招待
- 招待リンク発行 or メール指定


## 4. 画面要件

| ページ名 | 内容 |
|:--|:--|
| ログインページ | メール・パスワードでログイン |
| 新規登録ページ | ユーザー名、メール、パスワード登録 |
| チーム一覧ページ | 所属チーム一覧表示 |
| チーム作成ページ | チーム名入力フォーム |
| チーム詳細ページ | TODO一覧表示、新規追加フォーム |
| 招待ページ | メール入力、招待ボタン |


## 5. 非機能要件

- Docker Composeを利用して一括起動可能
- Clean Architectureベースの実装
- GORMを使用したDB操作
- テストの整備 (Usecase層中心)
- バリデーションチェック
- APIドキュメントはREADME等に記載

---

## 6. 使用技術

| 分類 | 技術 |
|:--|:--|
| バックエンド | Go (Gin)、GORM |
| フロントエンド | Next.js 14 (App Router)、Tailwind CSS |
| 認証 | JWT |
| インフラ | Docker / Docker Compose、MySQL |


## 7. データベース設計 (仮)

| テーブル名 | 説明 |
|:--|:--|
| customers | ユーザー情報 (ID, メール, パスワード, 名前) |
| teams | チーム情報 (ID, チーム名, 説明) |
| team_members | チームメンバー管理 (team_id, user_id,role) |
| todos | TODO情報 (ID, title, completed, team_id, user_id) |
