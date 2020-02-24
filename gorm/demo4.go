package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Article struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	CreateTime string `db:"create_time"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	ImageUrl   string `db:"imageurl"`
}

func main() {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", "postgres", "root", "go-web")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("connect pgsql failed, err:", err)
		return
	}
	db.LogMode(true)
	db.SingularTable(true)

	var article Article
	db.Where("id = ?", "1").Find(&article)
	fmt.Println("article:", article)
}
