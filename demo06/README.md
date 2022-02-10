### 传参的服务注册 && 服务发现均衡负载

go run register.go --server_address :8088
go run register.go --server_address :8089
go run register.go --server_address :8090


> go-micro的服务发现的算法由github.com/micro/go-micro/client/selector下的selector提供，目前提供了两种算法

+ RoundRobin(轮询算法)
+ Random(随机算法)