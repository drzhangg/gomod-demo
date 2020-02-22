package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type","alert")

		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %v, username is %v, password is %v, hobby is %v", type1, username, password, hobbys))
	})
	router.Run(":8080")
}
