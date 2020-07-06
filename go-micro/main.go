package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	router := gin.Default()
	router.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})

	router.Handle("GET", "/news", func(c *gin.Context) {
		c.String(200, "news api")
	})

	service := web.NewService(
		web.Address(":8001"),
		web.Handler(router),
	)

	service.Run()
}
