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

# run this dockerfile from toplevel directory
FROM golang:1.13.0-alpine as builder
LABEL maintainer "Yoshi Yamaguchi <@ymotongpoo>"
COPY ./server /server/
COPY ./genproto /genproto/
WORKDIR /server
RUN CGO_ENABLED=0 go build -o server

FROM busybox as runner
WORKDIR /workdir
COPY --from=builder /server/server /workdir/server
CMD ["/workdir/server"]