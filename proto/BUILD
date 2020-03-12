load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "fibonacci_proto",
    srcs = ["fib.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "fibonacci_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/ymotongpoo/bazel-kind-grpc-sample/proto",
    proto = ":fibonacci_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    embed = [":fibonacci_go_proto"],
    importpath = "github.com/ymotongpoo/bazel-kind-grpc-sample/proto",
    visibility = ["//visibility:public"],
)
