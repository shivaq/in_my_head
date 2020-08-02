############
# 1st stage
############

# Golang のオフィシャルイメージを使用
# OS は Alpine Linux
FROM golang:1.14-alpine AS build

# WORKDIR で、続くインストラクションが実行される
WORKDIR /src/
COPY main.go /src/

# build 時に実行されるコマンド
# CGO_ENABLED →C を使わない場合 0
RUN CGO_ENABLED=0 go build -o /bin/demo

############
# 2nd stage
############

# 空っぽイメージから空っぽコンテナを作成
FROM scratch

# 1st ステージの バイナリファイルをコピー
COPY --from=build /bin/demo /bin/demo
# 2nd ステージで作られたコンテナのみが、デプロイされる。
# よって、1st ステージで 350MB だったのが、2nd ステージのみだから 6MBになる


# ENTRYPOINT -> コンテナが実行ファイルとして動くようになる
# /bin/demo というコマンドが実行される
ENTRYPOINT ["/bin/demo"]