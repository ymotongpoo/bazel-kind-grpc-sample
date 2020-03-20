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