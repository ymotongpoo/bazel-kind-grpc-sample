module github.com/ymotongpoo/bazel-kind-grpc-sample/client

go 1.14

replace github.com/ymotongpoo/bazel-kind-grpc-sample/genproto => ../genproto

require (
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/ymotongpoo/bazel-kind-grpc-sample/genproto v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7 // indirect
	google.golang.org/grpc v1.23.1
)
