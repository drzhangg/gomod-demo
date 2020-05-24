package common

import "fmt"

func Arr() {
	a := [2]int{1, 2}
	b := []int{1, 2, 3, 4, 5, 6, 7}

	for _, i := range a {
		fmt.Printf("a 数组遍历：%d\n", i)
	}

	for _, i := range b {
		fmt.Printf("b 切片遍历：%d\n", i)
	}

	//输出数组的第二个元素
	fmt.Printf("a的第二个元素 %d\n", a[1])

	// 输出切片的第五个元素
	fmt.Printf("b的第五个元素 %d\n", b[4])
}
