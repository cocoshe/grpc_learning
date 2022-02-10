### micro 注册 consul 和服务发现

github.com/micro/go-micro/registry/consul
上面这个在1.14.0版本之后删除了，要用这个github.com/micro/go-plugins/registry/consul，或者换成etcd作为注册中心

```shell
go get github.com/micro/go-plugins/registry/consul

```