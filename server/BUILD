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
    importpath = "github.com/ymotongpoo/bazel-kind-grpc-sample/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    visibility = ["//:__pkg__"],
    embed = [":go_default_library"],
    cgo = False,
    out = "server",
)

container_image(
    name = "server_container",
    visibility = ["//:__pkg__"],
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
    visibility = ["//:__pkg__"],
    format = "Docker",
    image = ":server_container",
    registry = "gcr.io",
    repository = "$(project_id)/$(repo)",
    tag = "bazel",
)