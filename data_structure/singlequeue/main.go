package main

import (
	"errors"
	"fmt"
	"os"
)

/**
往队列里面添加数据，则rear的下标+1
从队列里面取数据，则front+1
*/
type Queue struct {
	maxSize int    //最大容量
	array   [4]int //数组容量
	front   int    //front 队列首部
	rear    int    //rear 队列尾部
}

//添加数据到队列
func (q *Queue) AddQueue(val int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("queue is full")
	}

	//添加数据，队列尾部下标进行自增
	q.rear++
	//将数据添加到队列中
	q.array[q.rear] = val

	return nil
}

//从队列中取出数据
func (q *Queue) GetQueue() (val int, err error) {
	//判断队列是否为空
	//当front的值和rear的值相等时，说明对列为空
	if q.front == q.rear {
		return -1, errors.New("queue is empty")
	}
	//当取出一个值时，front值要往前移
	q.front++
	val = q.array[q.front]
	return val,nil
}

//显示队列中的数据
func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

func main() {
	q := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {

		fmt.Println("1. 输入add，表示添加数据到队列")
		fmt.Println("2. 输入get，表示从队列获取数据")
		fmt.Println("3. 输入show，表示显示队列")
		fmt.Println("4. 输入exit，退出程序")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			if err := q.AddQueue(val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}

		case "get":

		case "show":
			q.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
