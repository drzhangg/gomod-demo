package main

import (
	"github.com/gin-gonic/gin"
	"gomod-demo/demo"
	"net/http"
)

const dateFormat = "2006-01-02 15:04:05"
const dateTimeFormat = "20060102150405"

func main() {
	router := gin.Default()
	router.GET("/test", Test)
	router.Run(":9090")
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

func Grpc(c *gin.Context) {
	demo.Service
}
