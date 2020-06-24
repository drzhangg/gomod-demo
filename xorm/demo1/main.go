package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

func main() {
	result := GetOne("1")
	fmt.Println(result.Id, result.Name, result.Id)

	r := GetOne1("2").(struct {
		Id   int
		Name string
		Age  int
	})
	fmt.Println(r.Id, r.Name, r.Age)
}

func GetOne(id string) struct {
	Id   int
	Name string
	Age  int
} {

	var user = struct {
		Id   int
		Name string
		Age  int
	}{}
	has, err := db().Table("user").Where("id = ?", id).Get(&user)
	if err != nil || !has {
		log.Printf("db get failed, err: %v\n", err)
	}
	return user
}

func GetOne1(id string) interface{} {

	var user = struct {
		Id   int
		Name string
		Age  int
	}{}
	has, err := db().Table("user").Where("id = ?", id).Get(&user)
	if err != nil || !has {
		log.Printf("db get failed, err: %v\n", err)
	}
	return user
}

func db() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Print(err)
	}
	return engine
}
