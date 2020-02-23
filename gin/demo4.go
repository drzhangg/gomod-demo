package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	User     string `json:"user" xml:"user" form:"user" uri:"user"`
	Password string `json:"password" xml:"password" form:"password" uri:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.User != "admin" || user.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user or password is failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
	})
	r.Run(":8080")
}
