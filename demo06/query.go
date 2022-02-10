package main

import (
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func main() {

	// 连接 consul
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	for true {
		services, err := cr.GetService("productService")
		if err != nil {
			log.Println(err)
		}

		// 使用selector 的 random随机获得一个实例
		next := selector.Random(services)
		svc, err := next()
		if err != nil {
			log.Println(err)
		}

		log.Println("[测试输出]: ", svc.Address)
		time.Sleep(time.Second)
	}

}
