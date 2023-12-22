build-server-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o target/server/game2048_linux-arm64 ./cmd/server

build-server-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o target/server/game2048_darwin-arm64 ./cmd/server

clean:
	rm -rf ./target