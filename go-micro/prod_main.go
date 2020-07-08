package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("47.103.9.218:8500"))

	router := gin.Default()
	router.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})

	router.Handle("GET", "/news", func(c *gin.Context) {
		c.String(200, "news api")
	})

	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(consulReg),
	)

	server.Run()
}
