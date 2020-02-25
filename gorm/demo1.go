package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Demo struct {
	Id     int64    `json:"id",db:"id,omitempty"`
	Name   []string `json:"name",db:"name,omitempty"`
	Age    int      `json:"age",db:"age,omitempty"`
	Salary int      `db:"salary,omitempty"`
}

var (
	N = "a_%s"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select name from people where id =1")
	if err != nil {
		fmt.Printf("Scan failed,err:%v", err)
	}

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			fmt.Printf("Scan failed,err:%v", err)
		}
	}
	fmt.Println(name)

	arr := strings.Split(name, ",")
	fmt.Println(arr)

	bytes, err := json.Marshal(&Demo{
		Id:     0,
		Name:   arr,
		Age:    0,
		Salary: 0,
	})
	if err != nil {
		fmt.Printf("json marshal failed,err:%v", err)
	}
	fmt.Println(string(bytes))

}
