# FROM
# DockerイメージのベースとなるイメージをFROMで指定
# Docker Hubで公開されているもの
FROM golang:1.9  

# RUN
# Dockerイメージビルド時にコンテナ内で実行するコマンド
RUN mkdir /echo

# COPY
# ホストマシン上のファイルやディレクトリをコンテナ内にコピーする
COPY main.go /echo

# CMD
# Dockerコンテナとして実行する際にコンテナ内で実行するプロセスを指定
CMD ["go", "run", "/echo/main.go"]

# ビルド
# コマンドは docker image build -t イメージ名[: タグ名] Dockerfile 配置ディレクトリのパス
# ex. docker image build -t example/echo:latest .
# コンテナの実行
# ex.  docker container run -d -p 9000:8080 example/echo:latest
# curlでGETリクエストを送信することでレスポンスを得る