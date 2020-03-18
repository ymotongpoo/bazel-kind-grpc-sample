load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_bundle")
load("@io_bazel_rules_docker//contrib:push-all.bzl", "container_push")

exports_files(["LICENSE"])

# gazelle:prefix github.com/ymotongpoo/bazel-kind-grpc-sample
gazelle(name = "gazelle")

container_bundle(
    name = "all_container",
    images = {
        "gcr.io/$(project_id)/$(repo)/client": "//client:client_container",
        "gcr.io/$(project_id)/$(repo)/server": "//server:server_container",
    },
)

container_push(
    name = "push_all",
    bundle = ":all_container",
    format = "Docker",
)