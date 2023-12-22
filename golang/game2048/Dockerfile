FROM ubuntu:latest
LABEL authors="akapralov"

WORKDIR /app

COPY target/server/game2048_linux-arm64 ./server

ENTRYPOINT ["/app/server"]
