package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro/prodservice"
	"net/http"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("47.103.9.218:8500"))

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.Handle("GET", "/prods", func(c *gin.Context) {
			c.JSON(http.StatusOK, prodservice.NewProdList(5))
		})
	}

	server := web.NewService(
		web.Name("prodservice1"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(consulReg),
	)

	server.Run()
}
