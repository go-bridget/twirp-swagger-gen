.PHONY: all build test

all: build test clean

build:
	go build -o build/ github.com/go-bridget/twirp-swagger-gen/cmd/...

test:
	./build/twirp-swagger-gen -in example/example.proto -out example/simple/example.swagger.json -host test.example.com
	./build/twirp-swagger-gen -in example/google_timestamp.proto -out example/simple/google_timestamp.swagger.json -host test.example.com

	# use go run so we do not have install buf command
	go run github.com/bufbuild/buf/cmd/buf@latest generate --template example/buf.gen.yaml --path example

clean:
	go fmt ./...
	go mod download
	go mod tidy

google:
	git clone https://github.com/googleapis/googleapis google
