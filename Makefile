all: build/myapp.wasm webserver

build/myapp.wasm: myapp.go
	GOOS=js GOARCH=wasm go build -o $@ $<

webserver: webserver.go
	go build $<

.PHONY: all