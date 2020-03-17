load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image",
    "container_push",
    "container_layer",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/ymotongpoo/bazel-kind-grpc-sample/client",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:go_default_library",
        "@com_github_labstack_echo//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "client",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
    cgo = False,
    out = "client",
)

container_image(
    name = "client_container",
    base = "@distroless_base_debian10_debug//image",
    directory = "/workdir",
    workdir = "/workdir",
    files = [
        ":client",
    ],
    cmd = [
        "/workdir/client"
    ]
)

container_push(
    name = "client_push",
    format = "Docker",
    image = ":client_container",
    registry = "gcr.io",
    repository = "$(PROJECT_ID)/$(REPO)",
    tag = "bazel",
)