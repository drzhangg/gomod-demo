package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

var Db *gorm.DB

//root:root@tcp(localhost:3306)/demo
const (
	username = "root"
	password = "root"
	network  = "tcp"
	server   = "47.103.9.218"
	port     = 3306
	database = "test1"
)

func init() {
	db, err := gorm.Open("mysql", "root:root123456@(47.103.9.218:3306)/test?charset=utf8")
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	//defer Db.Close()

	fmt.Println(db.DB().Ping())

	db.SingularTable(true)
	db.LogMode(true)

	Db = db
}

func main() {

	r := gin.Default()
	r.GET("/getUser", GetUser)

	r.Run(":8082")
}

func GetUser(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	user, err := GetUserByName(name)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUserByName(name string) (user User, err error) {
	//err = Db.Table("user").Where(fmt.Sprintf("name = '%s'",name)).First(&user).Error
	err = Db.Table("user").Where("name = ?", name).First(&user).Error
	return user, err
}

func GetUserByName2(name string) (user User, err error) {
	err = Db.Raw("select * from user where name = '" + name + "'").Scan(&user).Error
	return user, err
}

//Query
func QueryOne(db *gorm.DB) {
	var user User
	db.Raw("select name,age from user where name = ?", "jerry").Scan(&user)

	fmt.Println(user)
}

func Insert(db *gorm.DB) {
	sql := `insert into user (name,age) value (?,?)`
	db.Exec(sql, "tom", 30)
}
