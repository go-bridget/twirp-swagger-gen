# twirp-swagger-gen

A Twirp RPC Swagger/OpenAPI 2.0 generator 

# Usage

Installing the generator for protoc/buf:

```
go get -u github.com/candy-digital/twirp-swagger-gen
```

This should pull in the cmd/ folder as well. Installing just the binaries:

```
go install github.com/candy-digital/twirp-swagger-gen/cmd/...
```

Running the standalone version:

```
twirp-swagger-gen \
	-in example/example.proto \
	-out example/example.swagger.json \
	-host test.example.com
```

Running the protoc code with [buf.build](https://buf.build) (buf.gen.yaml):

```
version: v1
plugins:
  - name: twirp-swagger
    opt: hostname=api.example.com,path_prefix=/api/v1
    out: example/buf
```

Other? Try to figure it out, then open a PR for the README.

# Why?

The project
[elliots/protoc-gen-twirp_swagger](https://github.com/elliots/protoc-gen-twirp_swagger)
is [defunct due to upstream changes to grpc-ecosystem
dependencies](https://github.com/elliots/protoc-gen-twirp_swagger/issues/25).

This project is a rewrite, that relies on both the official OpenAPI
structures, and a generic .proto file parser. The output should be line
compatible - my goal was just to replace the generator with a working one
without still being exposed to breaking changes from the gRPC ecosystem
packages.

The generated output is suitable for Swagger-UI.
