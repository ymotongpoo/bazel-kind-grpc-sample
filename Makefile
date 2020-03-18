# Copyright 2019 Yoshi Yamaguchi
#
# Licensed under the Apache License, Version 2.0 (the "License");
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

.PHONY: clean gazelle proto ferver client run

clean:
	bazel clean --expunge

gazelle:
	bazel run :gazelle -- update-repos -from_file=client/go.mod -to_macro=repositories.bzl%go_repositories
	bazel run :gazelle -- update-repos -from_file=server/go.mod -to_macro=repositories.bzl%go_repositories

proto:
	bazel build //proto:go_default_library

server:
	bazel build //server:server

client:
	bazel build //client:client

push_all:
	bazel build :push_all

run:
	docker-compose up
