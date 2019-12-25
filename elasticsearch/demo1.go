package main

import (
	"github.com/go-log/log/log"
	"github.com/olivere/elastic"
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

func init() {
	log.New()
}

func main() {

}
