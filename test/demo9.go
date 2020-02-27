package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}

func main() {
	router := gin.Default()
	router.POST("/Login", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
			return
		}
		if login.User != "jerry" || login.Password != "123"{
			c.JSON(http.StatusOK, gin.H{"err": "user or password is failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": login})
	})



	router.Run(":8080")
}
