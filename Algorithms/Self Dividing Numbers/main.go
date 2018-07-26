package main

import "fmt"

func main() {

	fmt.Println(selfDividingNumbers(1,1000000000))
}

func selfDividingNumbers(left int, right int) []int {

	var result []int

	for i := left; i<=right; i++ {

		if i <= 9 {
			result = append(result, i)
		}else{
			j := i
			flag := 0
			for {
				x := j%10
				if x == 0 {
					flag = 1
					break
				}
				if i%x != 0 {
					flag = 1
					break
				}
				if j/10 < 1 {
					break
				}
				j = j/10

			}
			if flag == 0 {
				result = append(result, i)
			}
		}
	}
	return result
}
