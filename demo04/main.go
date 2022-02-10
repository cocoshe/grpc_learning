package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

func main() {
	//n 作为路由
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user api",
		})
	})

	server := web.NewService(
		web.Address(":8081"),
		web.Metadata(map[string]string{"protocol": "http"}),
		web.Handler(r))

	err := server.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
