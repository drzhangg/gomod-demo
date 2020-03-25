package main

import "fmt"

type HeroNode struct {
	id       int
	name     string
	nickname string
	next     *HeroNode //这里表示指向下一个节点
}

//给链表插入一个节点
func InsertHeroNode(head *HeroNode, newNode *HeroNode) {
	/*
		1.先找到这个链表的最后一个节点
		2.创建一个辅助节点
		3.将newNode加入链表的最后
	*/
	temp := head
	for {
		if temp.next == nil { //说明到了最后一个节点
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}

	temp.next = newNode
}

// 根据输入的链表编号，从小到大依次插入
func InsertNode(head *HeroNode, newNode *HeroNode) {
	temp := head
	flag := true

	for {
		if temp.next == nil { //说明是最后一个节点
			break
		} else if temp.next.id > newNode.id { //说明newNode应该在head的前面
			break
		} else if temp.next.id == newNode.id {
			flag = false
			break
		}

		temp = temp.next //将新节点的值赋到next上，以此新增
	}

	if !flag {
		fmt.Printf("对不起，该id已经存在:%d", newNode.id)
		return
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}
}

//显示链表的所有信息
func ListHeroNode(head *HeroNode) {
	/*
		1.先判断这个链表是否为空
		2.不断的遍历链表，获取最后的值
	*/
	tmp := head
	if tmp.next == nil {
		fmt.Println("链表为空")
		return
	}

	for {
		fmt.Printf("[%d, %s, %s ] --> ", tmp.next.id, tmp.next.name, tmp.next.nickname)
		//不断的查看后面节点是否为空
		tmp = tmp.next
		if tmp.next == nil {
			break
		}
	}
}

func main() {
	node := &HeroNode{}

	temp1 := &HeroNode{
		id:       1,
		name:     "jerry",
		nickname: "mouse",
	}
	temp2 := &HeroNode{
		id:       2,
		name:     "tom",
		nickname: "cat",
	}

	temp3 := &HeroNode{
		id:       3,
		name:     "houser",
		nickname: "女主人",
	}

	//InsertHeroNode(node, temp1)
	//InsertHeroNode(node, temp2)
	InsertNode(node,temp2)
	InsertNode(node,temp3)
	InsertNode(node,temp1)

	ListHeroNode(node)

}
