// Copyright 2019 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"

	fib "github.com/ymotongpoo/bazel-kind-grpc-sample/fib"
)

var (
	fibServiceAddr string
	fibServicePort string

	conn    *grpc.ClientConn
	timeout = 3 * time.Second

	buffer *sharedList
)

type sharedList struct {
	sync.Mutex
	list []int64
}

func (s *sharedList) Len() int64 {
	s.Lock()
	l := len(s.list)
	s.Unlock()
	return int64(l)
}

func (s *sharedList) Add(n int64) {
	s.Lock()
	s.list = append(s.list, n)
	s.Unlock()
}

func (s *sharedList) Nth(n int64) int64 {
	return s.list[n]
}

func init() {
	fibServiceAddr = os.Getenv("FIB_SERVICE_ADDR")
	fibServicePort = os.Getenv("FIB_SERVICE_PORT")

	if fibServiceAddr == "" || fibServicePort == "" {
		log.Fatal("make sure to specify environment variables: FIB_SERVICE_ADDR and FIB_SERVICE_PORT")
	}

	buffer = &sharedList{
		list: []int64{1, 1, 2},
	}
}

func main() {
	listener, err := net.Listen("tcp", fibServiceAddr+":"+fibServicePort)
	if err != nil {
		log.Fatalf("Failed to open listner: %v", err)
	}

	server := grpc.NewServer()
	service := new(fibonacciServiceServer)
	fib.RegisterFironacciServiceServer(server, service)
	err = server.Serve(listener)
	log.Fatalf("Error serving gRPC service: %v", err)
}

type fibonacciServiceServer struct{}

func (f *fibonacciServiceServer) GetNth(c context.Context, r *fib.GetFibonacciRequest) (*fib.GetNthResponse, error) {
	num := fibNth(r.GetNth())
	return &fib.GetNthResponse{
		Number: num,
	}, nil
}

func (f *fibonacciServiceServer) GetListUntilNth(c context.Context, r *fib.GetFibonacciRequest) (*fib.GetListUntilNthResponse, error) {
	n := r.GetNth()
	fibNth(n)
	return &fib.GetListUntilNthResponse{
		Numbers: buffer.list[:n-1],
	}, nil
}

func fibNth(n int64) int64 {
	if buffer.Len() < n {
		fibToN(n)
	}
	return buffer.Nth(n - 1)
}

func fibToN(n int64) {
	l := buffer.Len()
	x := buffer.list[l-2]
	y := buffer.list[l-1]
	for i := l; i < n; i++ {
		x, y = y, x+y
		buffer.Add(y)
	}
}
