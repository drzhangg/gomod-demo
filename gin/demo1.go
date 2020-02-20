package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		//var l Login
		//if err := c.ShouldBind(&l); err != nil {
		//	fmt.Println("shouldbind failed,err:", err)
		//	return
		//}
		user,_ := c.Get("user")

		fmt.Println("user:",c.Param("user"))
		fmt.Println("user2:",user.(string))
		fmt.Println("pwd:",c.Param("pwd"))

		if c.Param("user") == "zhang" && c.Param("pwd") == "123"  {
			c.JSON(http.StatusOK,gin.H{"msg":"login success"})
		}else {
			c.JSON(http.StatusOK,gin.H{"err":"user or pwd failed"})
		}

	})

	r.Run(":8080")
}
