package main

import (
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main() {
	// 1. 连接到 consul
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	// 2. 根据 service name 获取对应的微服务列表
	services, err := cr.GetService("productService")
	if err != nil {
		log.Println(err)
	}

	// 3. 使用 random 随机获取其中一个实例
	next := selector.Random(services)
	svc, err := next()
	if err != nil {
		log.Println(err)
	}

	log.Println("[测试输出]: ", svc.Id, svc.Address, svc.Metadata)

}
