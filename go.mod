module grpc_learning

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/hashicorp/consul/api v1.12.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1 // indirect
	github.com/micro/micro v1.18.0 // indirect
	github.com/micro/micro/v3 v3.9.0 // indirect
	google.golang.org/grpc v1.44.0
	google.golang.org/grpc/examples v0.0.0-20220209221444-a354b1eec350 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
