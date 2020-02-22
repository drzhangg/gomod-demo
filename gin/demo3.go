package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/form", func(c *gin.Context) {
		//	type1 := c.DefaultPostForm("type","alert")
		//
		//	username := c.PostForm("username")
		//	password := c.PostForm("password")
		//	hobbys := c.PostFormArray("hobby")
		//	c.String(http.StatusOK, fmt.Sprintf("type is %v, username is %v, password is %v, hobby is %v", type1, username, password, hobbys))

		//file, _ := c.FormFile("file")
		//
		//if err := c.SaveUploadedFile(file, "/Users/drzhang"); err != nil {
		//	fmt.Sprintf("save failed,err :", err.Error())
		//}

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadGateway, fmt.Sprintf("get err %v", err.Error()))
		}
		files := form.File["file"]
		for _, file := range files {
			bytes, err := ioutil.ReadFile(file.Filename)
			if err != nil {
				fmt.Println("read file failed, err:", err)
			}
			filename := "/Users/drzhang/"
			err = ioutil.WriteFile(filename, bytes, 777)
			if err != nil {
				fmt.Println("write file failed, err:", err)
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("file is %v", len(files)))

	})
	router.Run(":8080")
}
