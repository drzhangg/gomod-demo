package main

import "fmt"

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
	for _, v := range chessMap {
		for _, k := range v {
			fmt.Printf("%d  ", k)
		}
		fmt.Println()
	}

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

	for i, v := range nodes {
		fmt.Printf("%d：%d  %d  %d\n", i, v.row, v.low, v.val)
	}

	fmt.Println(nodes)
}
