package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func main() {
	cr := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user api",
		})
	})

	srv := web.NewService(
		web.Name("productService"),
		web.Registry(cr),
		web.Address(":8081"),
		web.Metadata(map[string]string{"protocol": "http"}),
		web.Handler(r),
	)

	// 命令行设置参数
	srv.Init()

	_ = srv.Run()

}
