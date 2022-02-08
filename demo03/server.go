package main

import (
	"context"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"grpc_learning/demo03/pb"
	"log"
	"net"
)

// 定义类
type Children struct {
}

// 绑定类方法，实现接口
func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello " + p.Name
	return p, nil
}

func main() {
	// 把grpc服务，注册到consul上
	//1. 初始化consul 配置
	consulConfig := api.DefaultConfig()

	//2. 创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Println(err)
		return
	}
	//3. 告诉consul， 即将注册的服务配置信息
	reg := api.AgentServiceRegistration{
		ID:      "MyService",
		Name:    "grpc and consul",
		Tags:    []string{"grpc", "consul"},
		Port:    8800,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}

	//4. 注册grpc到consul上
	consulClient.Agent().ServiceRegister(&reg)
	/**********以下为grpc远程调用***************/
	// 1. 初始化 grpc 对象
	grpcServer := grpc.NewServer()

	// 2. 注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	// 3. 设置监听
	//listener, err := net.Listen("tcp", "192.168.2.39:8800")
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Println("Listen err: ", err)
		return
	}
	defer listener.Close()

	// 4. 启动服务
	grpcServer.Serve(listener)
}
