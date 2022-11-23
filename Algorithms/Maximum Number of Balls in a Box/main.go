package main

import "fmt"

func main() {
	fmt.Println(countBalls(19, 28))
}

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	lastCount := calBoxNumber(lowLimit)
	m[lastCount]++
	for lowLimit = lowLimit + 1; lowLimit <= highLimit; lowLimit++ {
		c := suffix9Count(lowLimit - 1)
		if c != 0 {
			lastCount = lastCount - 9*c + 1
		} else {
			lastCount = lastCount + 1
		}
		m[lastCount]++
	}
	ans := 0
	for _, v := range m {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func calBoxNumber(b int) int {
	sum := 0
	for b > 0 {
		sum += b % 10
		b = b / 10
	}
	return sum
}

func suffix9Count(n int) int {
	c := 0
	for {
		if n%10 != 9 {
			break
		}
		c++
		n = n / 10
	}
	return c
}
