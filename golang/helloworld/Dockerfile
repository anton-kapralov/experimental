FROM ubuntu:latest
LABEL authors="akapralov"

WORKDIR /app

COPY target/linux/helloworld_arm64 ./helloworld

ENTRYPOINT ["/app/helloworld"]
