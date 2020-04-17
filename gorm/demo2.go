package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

//root:root@tcp(localhost:3306)/demo
//root:root@tcp(47.103.9.218:3306)/test1?charset=utf8
const (
	username = "root"
	password = "root"
	network  = "tcp"
	server   = "47.103.9.218"
	port     = 3306
	database = "test1"
)

func main() {
	//db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database))

	db, err := gorm.Open("mysql", "root:root@(47.103.9.218:3306)/test1?charset=utf8")
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	defer db.Close()

	fmt.Println(db.DB().Ping())

	db.SingularTable(true)
	db.LogMode(true)

	QueryOne(db)
	Insert(db)
}

//Query
func QueryOne(db *gorm.DB) {
	var user User
	//db.Find(&user)
	db.Raw("select name,age from user where name = ?", "jerry").Scan(&user)

	fmt.Println(user)
}

func Insert(db *gorm.DB) {
	sql := `insert into user (name,age) value (?,?)`
	db.Exec(sql, "tom", 30)
}
