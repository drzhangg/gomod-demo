package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Demo struct {
	Id     int64  `json:"id",db:"id,omitempty"`
	Name   string `json:"name",db:"name,omitempty"`
	Age    int    `json:"age",db:"age,omitempty"`
	Salary int    `db:"salary,omitempty"`
}

var (
	N = "a_%s"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	if err != nil {
		fmt.Println(err)
	}

	demo := new(Demo)
	rows, err := db.Query("select * from people")
	if err != nil {
		fmt.Printf("Scan failed,err:%v", err)
	}

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}



	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//err = rows.Scan(&demo.Name, &demo.Age, &demo.Salary, &demo.Id)
	//if err != nil {
	//	fmt.Printf("Scan failed,err:%v", err)
	//}
	//fmt.Println(*demo)

	for rows.Next() {
		//err = rows.Scan(&demo.Name, &demo.Age, &demo.Salary, &demo.Id)
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
		}

		var value string
		for i,col := range values{
			if col == nil{
				value = "null"
			}else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}



		fmt.Println(*demo)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
