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

.PHONY: all clean server client run

all: protoc run

clean:
	docker rmi fib-client:1.13.0 -f
	docker rmi fib-server:1.13.0 -f

protoc: fib.proto
	protoc --go_out=plugins=grpc:genproto -I . ./fib.proto

server: server/main.go
	docker-compose build server

client: client/main.go
	docker-compose build client

run:
	docker-compose up
