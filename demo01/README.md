```shell
protoc --go_out=./grpc-srv/services ./profiles/Prod.proto
protoc --go-grpc_out=./grpc-srv/services ./profiles/Prod.proto
```
two param:
1. tgt dir
2. src proto file