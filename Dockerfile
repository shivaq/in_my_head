############
# 1st stage
############

# Golang のオフィシャルイメージを使用
# OS は Alpine Linux
FROM golang:1.14-alpine AS build

# for go mod download
RUN apk add --update --no-cache ca-certificates git

# WORKDIR で、続くインストラクションが実行される
WORKDIR /src/
COPY go.mod .
# COPY go.sum .

RUN go mod download
# COPY the source code as the last step
COPY . .

RUN ls -la
RUN pwd

# build 時に実行されるコマンド
# CGO_ENABLED →C を使わない場合 0
RUN CGO_ENABLED=0 go build -o /bin/main

############
# 2nd stage
############

# 空っぽイメージから空っぽコンテナを作成
FROM scratch

# 1st ステージの バイナリファイルをコピー
COPY --from=build /bin/main /bin/main
# 2nd ステージで作られたコンテナのみが、デプロイされる。
# よって、1st ステージで 350MB だったのが、2nd ステージのみだから 6MBになる


# ENTRYPOINT -> コンテナが実行ファイルとして動くようになる
# /bin/main というコマンドが実行される
ENTRYPOINT ["/bin/main"]