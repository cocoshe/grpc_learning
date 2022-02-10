package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func main() {
	// 添加 consul 地址
	cr := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	// 使用 gin 作为路由
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user api",
		})
	})

	// 初始化go micro
	srv := web.NewService(
		web.Name("productService"),
		web.Registry(cr),
		web.Address(":8081"),
		web.Metadata(map[string]string{"protocol": "http"}),
		web.Handler(r),
		web.Id("this is service id"),
	)

	_ = srv.Run()
}
