### Consul and grpc
通过consul来做服务发现

运行默认配置的consul
```shell
consul agent -dev
```
```go
go run server.go
go run client.go
```