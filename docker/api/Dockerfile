# 必須 - 基になるイメージ
FROM golang:alpine

# 環境変数の設定
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# 作業ディレクトリをホットロード環境に変更
WORKDIR /go/src/hot_reload_docker
ENV GO111MODULE=on

# 実行コマンド
RUN apk update

# コンテナが接続用にリッスンするポートの指定
EXPOSE 8080

# コンテナ内でホットリロード環境でmain.goを起動
RUN go get github.com/pilu/fresh
CMD ["fresh"]