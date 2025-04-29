# ベースイメージを指定
FROM node:23-alpine

# 作業ディレクトリを設定
WORKDIR /app

# キャッシュ利用で効率化するために別でコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install

# ソースコードをコピー
COPY . .

# ポート番号を指定
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "run", "dev"]