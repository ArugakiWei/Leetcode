package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(101001))
}

// n & (n-1) 可以消去 n 的 二进制表示的最后一个 1
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
