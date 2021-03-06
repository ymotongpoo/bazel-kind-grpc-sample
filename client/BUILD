# Copyright 2020 Yoshi Yamaguchi
#
# Licensed under the Apache License, Version 2.0 (the "License")
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image",
    "container_push",
    "container_layer",
)
load("@io_bazel_rules_k8s//k8s:object.bzl", "k8s_object")

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
    visibility = ["//:__pkg__"],
    embed = [":go_default_library"],
    cgo = False,
    out = "client",
)

container_image(
    name = "container",
    visibility = ["//:__pkg__"],
    base = "@distroless_base_debian10//image",
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
    name = "push",
    visibility = ["//:__pkg__"],
    format = "Docker",
    image = ":client_container",
    registry = "gcr.io",
    repository = "{PROJECT_ID}/{REPO}",
    tag = "bazel",
)

k8s_object(
    name = "gke_deploy",
    kind = "deployment",
    visibility = ["//:__pkg__"],
    images = {
        "gcr.io/{PROJECT_ID}/{REPO}/client": ":container"
    },
    cluster = "{K8S_CLUSTER_NAME}",
    template = "//manifests:client.gke.yaml",
    resolver_args = ["--allow_unused_images"],
)

