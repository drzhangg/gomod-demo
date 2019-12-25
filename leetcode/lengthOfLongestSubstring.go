package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	cur, max := 0, 1
	array := []rune(s)

	for x, v := range array {
		y := cur

		for x > y {
			if array[y] == v { //第几个字节等于v
				break
			}
			y++
		}
		if x == y {
			newMax := x - cur + 1
			if newMax > max {
				max = newMax
			}
		} else {
			cur = y + 1
		}
	}
	return max
}

func main() {

}
