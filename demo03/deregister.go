package main

import "github.com/hashicorp/consul/api"

func main() {
	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul 对象
	consulClient, _ := api.NewClient(consulConfig)

	// 3. 注销服务
	consulClient.Agent().ServiceDeregister("MyService")

}
