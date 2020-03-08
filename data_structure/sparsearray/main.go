package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	row int //行
	low int //列
	val int //值
}

func main() {

	//定义一个二维数组
	var chessMap [11][11]int
	chessMap[1][2] = 2
	chessMap[2][3] = 1

	//打印二维数组
	//for _, v := range chessMap {
	//	for _, k := range v {
	//		fmt.Printf("%d  ", k)
	//	}
	//	fmt.Println()
	//}

	var nodes []Node
	var node Node

	for i, v := range chessMap {
		for j, k := range v {
			if k != 0 {
				node = Node{
					row: i,
					low: j,
					val: k,
				}
				nodes = append(nodes, node)
			}
		}
	}
	//fmt.Println(len(nodes))

	filename := "chessmap.data"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, v := range nodes {

		bytes := []byte(fmt.Sprintf("%v", v))
		for i := 0; i < len(nodes); i++ {
			writer.Write(bytes)
			writer.WriteString("\n")
		}
		writer.Flush()
		//_ = ioutil.WriteFile(filename, bytes, 0666)
	}

	//将稀疏数组转成二维数组
	var chessMap2 [11][11]int

	for _, v := range nodes {
		//fmt.Println(i,v)
		chessMap2[v.row][v.low] = v.val
	}
	//chessMap2 = chessMap2
	for _, i := range chessMap2 {
		for _, j := range i {
			fmt.Printf("%d   ", j)
		}
		fmt.Println()
	}

	fmt.Printf("%v", 10*time.Second)

	fmt.Println(strings.ToLower("Abcd"))
	//fmt.Println(chessMap2)
}
