package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `db:"id"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
}

//root:root@tcp(localhost:3306)/demo
const (
	username = "root"
	password = "root"
	network  = "tcp"
	server   = "localhost"
	port     = 3306
	database = "demo"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database))
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	defer db.Close()

	QueryOne(db)
	Querys(db)
	Insert(db)
	Update(db)
	delete(db)
}

//Query
func QueryOne(db *sql.DB) {
	user := new(User)
	var sql = `select * from user where id = ?`
	rows := db.QueryRow(sql, 1)

	if err := rows.Scan(&user.Id, &user.UserName, &user.PassWord); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}

//Querys
func Querys(db *sql.DB) {
	user := new(User)
	users := []User{}
	rows, err := db.Query("select * from user ")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.PassWord, &user.UserName)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(*user)
		users = append(users, *user)
	}
	fmt.Println(users)
}

//insert
func Insert(db *sql.DB) {
	var sql = `insert into user(username,password) values ('%s','%s')`
	sql = fmt.Sprintf(sql, "zhang", "123456")
	fmt.Println(sql)
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("lastInserId:", lastInsertId)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rowsAffected:", rowsAffected)
}

//update
func Update(db *sql.DB) {
	var sql = `update user set password = '%s' where id = %d`
	sql = fmt.Sprintf(sql, "root", 1)
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//delete
func delete(db *sql.DB) {
	var sql = `delete from user where username='%s'`
	sql = fmt.Sprintf(sql, "zhang")
	result, err := db.Exec(sql)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
