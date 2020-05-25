package linklist

import (
	"fmt"
	"log"
)

// Item可以理解为泛型，也就是任意的类型
type Item interface {
}

// 一个节点，除了自身的数据之外，还必须指向下一个节点，尾部节点指向为nil
type LinkNode struct {
	Payload Item //Payload为任意数据类型
	Next    *LinkNode
}

func (head *LinkNode) Add(payload Item) {
	//这里采用尾部插入的方式，给链表添加元素
	point := head

	for point.Next != nil {
		point = point.Next
	}
	newNode := LinkNode{
		Payload: payload,
		Next:    nil,
	}
	point.Next = &newNode

	//头部插入
	//newNode := LinkNode{payload,nil}
	//newNode.Next = head
}

func NewLinkNode(length int) *LinkNode {
	if length <= 0 {
		fmt.Printf("链表长度必须大于0")
		log.Panic("链表长度必须大于0")
	}
	var head *LinkNode

	head = &LinkNode{}

	for i := 0; i < length; i++ {
		var newNode *LinkNode
		newNode = &LinkNode{Payload: i}
		newNode.Next = head
		head = newNode
	}
	return head
}
