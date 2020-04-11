package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
)

var client *elastic.Client

var host = "http://47.103.9.218:9200"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

type Class struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func init() {
	file := "./log.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println("os.OpenFile filed,err: ", err)
		return
	}

	cfg := []elastic.ClientOptionFunc{
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
		elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)),
	}

	client, err = elastic.NewClient(cfg...)
	if err != nil {
		fmt.Println("elastic.NewClient filed,err: ", err)
		return
	}
	return
}

func main() {
	id := "1"
	exist := isExists(id)
	fmt.Println(exist)

	c := Class{}
	c.Get(id)
	fmt.Println(c.Name, c.Age, c.Score)
}

func isExists(id string) bool {
	defer client.Stop()
	exist, _ := client.Exists().Index("student").Type("class").Id(id).Do(context.Background())
	if !exist {
		log.Println("ID may be incorrect! ", id)
		return false
	}
	return true
}

func (class *Class) Get(id string) {
	getResult, err := client.Get().Index("student").Type("class").Id(id).Do(context.Background())
	if err != nil {
		fmt.Println("client get failed,err: ", err)
		return
	}

	bytes, _ := getResult.Source.MarshalJSON()
	json.Unmarshal(bytes, &class)

}

type SearchStruct struct {
	Category   string `json:"category"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Title      string `json:"title"`
	UpdateTime int64  `json:"update_timestamp"`
	Order      string `json:"order"`
	Sort       string `json:"sort"`
}

func Search(ss *SearchStruct) (class []Class) {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("category",ss.Category))
	query = query.Must(elastic.NewMatchQuery("title"),ss.Title)
}
