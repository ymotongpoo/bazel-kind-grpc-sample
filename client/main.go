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

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/fib.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/labstack/echo"
	fib "github.com/ymotongpoo/bazel-kind-grpc-sample/proto"
)

var (
	fibServiceAddr string
	fibServicePort string
	httpPort       string

	conn    *grpc.ClientConn
	timeout = 3 * time.Second
)

func init() {
	fibServiceAddr = os.Getenv("FIB_SERVICE_ADDR")
	fibServicePort = os.Getenv("FIB_SERVICE_PORT")
	httpPort = os.Getenv("HTTP_PORT")

	if fibServiceAddr == "" || fibServicePort == "" || httpPort == "" {
		log.Fatal("make sure to specify environment variables: FIB_SERVICE_ADDR, FIB_SERVICE_PORT and HTTP_PORT")
	}
}

func main() {
	ctx := context.Background()
	var err error
	conn, err = grpc.DialContext(ctx, fibServiceAddr+":"+fibServicePort,
		grpc.WithInsecure(),
		grpc.WithTimeout(timeout))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	e := echo.New()
	e.GET("/", helloHandler)
	e.GET("/nth/:num", nthHandler)
	e.GET("/nthlist/:num", nthListHandler)
	e.Logger.Fatal(e.Start("0.0.0.0:" + httpPort))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello, fib")
}

func nthHandler(c echo.Context) error {
	ctx := context.Background()
	svc := fib.NewFironacciServiceClient(conn)
	numStr := c.Param("num")
	log.Printf("request: %v\n", numStr)
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("parsed: %vth\n", numStr)
	nthReq := &fib.GetFibonacciRequest{
		Nth: num,
	}
	resp, err := svc.GetNth(ctx, nthReq)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("%dth fib number is %d", num, resp.GetNumber()))
}

func nthListHandler(c echo.Context) error {
	ctx := context.Background()
	svc := fib.NewFironacciServiceClient(conn)
	numStr := c.Param("num")
	log.Printf("request: %vth\n", numStr)
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Println(err)
		return err
	}
	nthReq := &fib.GetFibonacciRequest{
		Nth: num,
	}
	resp, err := svc.GetListUntilNth(ctx, nthReq)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("fib numbers until %dth are %v", num, resp.GetNumbers()))
}
