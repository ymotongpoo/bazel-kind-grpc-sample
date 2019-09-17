# docker-compose and gRPC sample

This is a sample application that demonstrates:

* docker-compose
  * docker multi stage build
  * build Go applications with docker-compose
  * communication between containers 
* gRPC
  * `protoc` for Go applications
  * simple gRPC server and client implementations

## prerequisites

* [Docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/)
  * make sure to have Docker 18.06.0+
* [Go](https://golang.org/)
* [gRPC](https://godoc.org/google.golang.org/grpc) (optional)
* [protoc-gen-go](https://github.com/golang/protobuf/protoc-gen-go) (optional)
* make (optional)

## How to run

### 0. (optional) Compile `.proto` file to Go code

Though this repository is already shipped with pre-compiled Go code from `.proto` file,
you can try compiling it with `protoc` and `protoc-gen-go`.

Given you have Go in your development environment, install gRPC and `protoc-gen-go`:

```
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ export PATH=$PATH:$GOPATH/bin
$ protoc --go_out=plugins=grpc:genproto -I . ./fib.proto
```

### 1. Run docker-compose

`docker-compose` build container images and launch those containers in the requiured order.

```
$ docker-compose up
```

Then you will see the log of client if it runs container successfully:

```
Starting fib_server ... done
Starting docker-compose-grpc-sample_client_1 ... done
Attaching to fib_server, docker-compose-grpc-sample_client_1
client_1  |
client_1  |    ____    __
client_1  |   / __/___/ /  ___
client_1  |  / _// __/ _ \/ _ \
client_1  | /___/\__/_//_/\___/ v3.3.10-dev
client_1  | High performance, minimalist Go web framework
client_1  | https://echo.labstack.com
client_1  | ____________________________________O/_______
client_1  |                                     O\
client_1  | ⇨ http server started on [::]:9999
```

### 2. Make requests to web client

You can make requests to the client by cURL, and you can confirm fibonacci numbers as response.

```
$ curl http://localhost:9999/nth/33
33th fib number is 3524578

$ curl http://localhost:9999/nthlist/20
fib numbers until 20th are [1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181]
```
