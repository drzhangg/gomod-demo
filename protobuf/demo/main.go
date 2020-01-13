package main

import (
	"context"
	"github.com/gin-gonic/gin"
	cus "gomod-demo/protobuf/pb/customer"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"
const dateTimeFormat = "20060102150405"

func main() {
	router := gin.Default()
	router.GET("/test", Test)
	router.GET("/grpc", Grpc)
	router.Run(":9090")
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

func Grpc(c *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	rpc := cus.NewCustomerClient(conn)
	reqBody1 := &cus.Id{Id: 1}
	res1, err := rpc.GetUser(context.Background(), reqBody1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from server：", res1.Name)

	reqBody2 := &cus.Name{Name: "jerry"}
	res2, err := rpc.GetActivity(context.Background(), reqBody2)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from server2:", res2.Name)
	c.String(http.StatusOK, "get grpc success,server1 :"+timeFormat(res1.Time, dateFormat))

}

func timeFormat(timeStamp int64, format ...string) string {
	t := time.Unix(timeStamp/1e3, 0)
	defaultFormat := dateTimeFormat
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return t.Format(defaultFormat)
}
