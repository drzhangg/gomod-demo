package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var  salay float32
	var ispass string

	fmt.Println("请输入名字")
	//time.Sleep(20)
	fmt.Scanln(&name)    //使用指针  改变name的值

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)    //使用指针  改变age的值

	fmt.Println("请输入薪水")
	fmt.Scanln(&salay)    //使用指针  改变name的值

	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&ispass)    //使用指针  改变name的值

	fmt.Println(name,age,salay,ispass)

}