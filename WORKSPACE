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

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.1/rules_go-v0.22.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.1/rules_go-v0.22.1.tar.gz",
    ],
    sha256 = "e6a6c016b0663e06fa5fccf1cd8152eab8aa8180c583ec20c872f4f9953a7ac5",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
    ],
    sha256 = "d8c45ee70ec39a57e7a05e5027c32b1576cc7f16d9dd37135b0eddde45cf1b10",
)

load("@io_bazel_rules_go//go:deps.bzl",
    "go_rules_dependencies",
    "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.13.8",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_protobuf",
    remote = "https://github.com/protocolbuffers/protobuf",
    commit = "d0bfd5221182da1a7cc280f3337b5e41a89539cf",
    shallow_since = "1581711200 -0800",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "rules_proto",
    sha256 = "2490dca4f249b8a9a3ab07bd1ba6eca085aaf8e45a734af92aad0c42d9dc7aaf",
    strip_prefix = "rules_proto-218ffa7dfa5408492dc86c01ee637614f8695c45",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/218ffa7dfa5408492dc86c01ee637614f8695c45.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/218ffa7dfa5408492dc86c01ee637614f8695c45.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# To generate repositories.bzl, run the following gazelle command.
# bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=repositories.bzl%go_repositories
load("//:repositories.bzl", "go_repositories")

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()

# Download the rules_docker repository
RULES_DOCKER_VER = "0.14.1"
RULES_DOCKER_HASH = "dc97fccceacd4c6be14e800b2a00693d5e8d07f69ee187babfd04a80a9f8e250"

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = RULES_DOCKER_HASH,
    strip_prefix = "rules_docker-%s" % RULES_DOCKER_VER,
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v%s/rules_docker-v%s.tar.gz" % (RULES_DOCKER_VER, RULES_DOCKER_VER)],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

# This is NOT needed when going through the language lang_image
# "repositories" function(s).
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "distroless_base_debian10",
    registry = "gcr.io",
    repository = "distroless/base-debian10",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    # Find the SHA256 digest value from the detials page of prebuilt containers.
    # https://console.cloud.google.com/gcr/images/distroless/GLOBAL/base-debian10
    digest = "sha256:732acc54362badaa64d9c01619020cf96ce240b97e2f1390d2a44cc22b9ba6a3",
)

# for debug
container_pull(
    name = "distroless_base_debian10_debug",
    registry = "gcr.io",
    repository = "distroless/base-debian10",
    tag = "debug",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    # Find the SHA256 digest value from the detials page of prebuilt containers.
    # https://console.cloud.google.com/gcr/images/distroless/GLOBAL/base-debian10
    digest = "sha256:8ca4526452afe5d03f53c41c76c4ddb079734eb99913aff7069bfd0d72457726",
)


# This requires rules_docker to be fully instantiated before
# it is pulled in.
# Download the rules_k8s repository
RULES_K8S_VER="0.4"
RULES_K8S_HASH="d91aeb17bbc619e649f8d32b65d9a8327e5404f451be196990e13f5b7e2d17bb"

http_archive(
    name = "io_bazel_rules_k8s",
    sha256 = RULES_K8S_HASH,
    strip_prefix = "rules_k8s-%s" % RULES_K8S_VER,
    urls = ["https://github.com/bazelbuild/rules_k8s/releases/download/v%s/rules_k8s-v%s.tar.gz" % (RULES_K8S_VER, RULES_K8S_VER)],
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories", "k8s_defaults")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s_go_deps.bzl", k8s_go_deps = "deps")

k8s_go_deps()

