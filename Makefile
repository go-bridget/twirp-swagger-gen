.PHONY: all build test

all: build test clean

build:
	go build -o build/ github.com/go-bridget/twirp-swagger-gen/cmd/...

test:
	./build/twirp-swagger-gen -in example/example.proto -out example/example.swagger.json -host test.example.com -I include
	./build/twirp-swagger-gen -in example/google_timestamp.proto -out example/google_timestamp.swagger.json -host test.example.com -I include

clean:
	go fmt ./...
	go mod download
	go mod tidy

google:
	git clone https://github.com/googleapis/googleapis google
