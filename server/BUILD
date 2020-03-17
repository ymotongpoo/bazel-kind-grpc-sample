load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/ymotongpoo/bazel-kind-grpc-sample/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
    cgo = False,
    out = "server",
)

container_image(
    name = "server_container",
    base = "@distroless_base_debian10_debug//image",
    directory = "/workdir",
    workdir = "/workdir",
    files = [
        ":server",
    ],
    cmd = [
        "/workdir/server"
    ]
)

container_push(
    name = "server_push",
    format = "Docker",
    image = ":server_container",
    registry = "gcr.io",
    repository = "$(PROJECT_ID)/$(REPO)",
    tag = "bazel",
)