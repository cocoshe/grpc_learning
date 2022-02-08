package main

import (
	"context"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"grpc_learning/demo03/pb"
	"log"
	"strconv"
)

func main() {
	// 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 创建 consul 对象 --(可以重新指定 consul 属性：IP/Port， 也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Println(err)
		return
	}

	// 服务发现,从 consul 上获取健康的服务
	services, _, err := consulClient.Health().Service("grpc and consul", "grpc", true, nil)
	if err != nil {
		log.Println(err)
		return
	}
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	/**********以下为grpc远程调用***************/
	// 1. 连接服务
	//conn, _ := grpc.Dial("192.168.2.39:8800", grpc.WithInsecure())
	//conn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())

	// 使用 服务发现 consul 上的 IP/Port 来与服务建立连接
	conn, _ := grpc.Dial(addr, grpc.WithInsecure())

	// 2. 初始化grpc客户端
	grpcClient := pb.NewHelloClient(conn)

	var person pb.Person
	person.Name = "cocoshe"
	person.Age = 18

	// 3. 调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &person)
	log.Println(p, err)

}
