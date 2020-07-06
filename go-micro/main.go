package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

func main() {
	etcdRegistry := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	router := gin.Default()
	router.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})

	router.Handle("GET", "/news", func(c *gin.Context) {
		c.String(200, "news api")
	})

	service := web.NewService(
		web.Name("product"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(etcdRegistry),
	)

	service.Run()
}
