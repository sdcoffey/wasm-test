bootstrap:
	brew install http-server

build:
	GOOS=js GOARCH=wasm go build -v -o go.wasm

run: build
	http-server -c-1
