# grpc-rest-example

[gRPC](https://grpc.io/) server/client example with **REST** gateway (see https://github.com/grpc-ecosystem/grpc-gateway) written in *Go*.

This is a simple *echo* server implementation that returns the requested message.


## Building the project

### Generating protocol buffer sources (docker)

If you have a running [docker](https://www.docker.com/) system you can use the supplied [protoc-gen.sh](https://github.com/bamling/grpc-rest-example/blob/master/scripts/protoc-gen.sh) script to generate the required [protocol buffers](https://developers.google.com/protocol-buffers/) source files.

```bash
./scripts/protoc-gen.sh
```

This should update/generate the required source files:

```bash
pkg/api/service.pb.go
pkg/api/service.pb.gw.go
```

### Generating protocol buffer sources (protoc)

Alternatively to the [docker](https://www.docker.com/) way, you can install `protoc` on your system (see [protocol buffers docs](https://developers.google.com/protocol-buffers/docs/downloads)) and generate the sources locally.

The following dependencies are required:

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

These enable `protoc` to generate *Go* sources for the **gRPC** gateway and provide us with a [swagger](https://swagger.io/) file on top of that.

Now you should be able to run the following command:


```bash
protoc \
    -I . \
    -I /usr/local/include \
    -I ${GOPATH}/src \
    -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=api/proto \
    --go_out=plugins=grpc:pkg/api \
    --grpc-gateway_out=logtostderr=true:pkg/api \
    --swagger_out=logtostderr=true:api/swagger \
    service.proto
``` 

### Building the example

The project utilises [Go Modules](https://github.com/golang/go/wiki/Modules) for the versioning of dependencies, thus *Go* in version >1.11 is required to reliably build the project.

Also ensure that *Modules* are enabled:

```bash
echo GO111MODULE=on
```

#### Building the server

```bash
go build ./cmd/server
```

#### Building the client

```bash
go build ./cmd/grpc-client
```


## Running the example

To start the `server`, execute the following:

```bash
./server
```

By default the **gRPC** server runs on port *8082* and the **HTTP** server/REST gateway on port *8080*.

Custom bind ports can be set via command line parameters:

```bash
./server -grpc-port 9092 -http-port 9090
```

The `grpc-client` automatically sends the message `Hello world!` to the **gRPC** server running on `localhost:8082`.

```bash
./grpc-client
{"level":"info","msg":"Response: echo.EchoMessage{Value:\"Hello world!\"}","time":"2019-04-03T14:43:40+02:00"}
```

To send different messages to different ports/addresses use:

```bash
./grpc-client -server-addr localhost:9092 -echo-message "Hello gRPC\!"
{"level":"info","msg":"Response: echo.EchoMessage{Value:\"Hello gRPC!\"}","time":"2019-04-03T14:45:26+02:00"}
```

Alternatively we can use *cURL* to query the **REST** endpoint:

```bash
curl -X POST http://localhost:8080/echo -d '{"value":"Hello REST gateway!"}'
{"value":"Hello REST gateway!"}
```