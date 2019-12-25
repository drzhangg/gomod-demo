package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}

	cur := head
	carry := 0

	if l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		val := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

func main() {
	fmt.Println(807 % 10)
}
