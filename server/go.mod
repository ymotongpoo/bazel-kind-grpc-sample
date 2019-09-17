module github.com/ymotongpoo/docker-compose-grpc-sample/server

go 1.12

replace github.com/ymotongpoo/docker-compose-grpc-sample/genproto => ../genproto

require (
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/ymotongpoo/docker-compose-grpc-sample/genproto v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/grpc v1.23.1
)
