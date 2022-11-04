package main

import "fmt"

func main() {
	fmt.Println(reachNumber(10000))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}

func reachNumber1(target int) int {
	dp := make(map[int]int)
	dp[0], dp[1], dp[-1] = 0, 1, 1
	mLastIndex := map[int]struct{}{
		-1: {},
		1:  {},
	}
	for i := 2; ; i++ {
		fmt.Println(len(mLastIndex))
		if v, ok := dp[target]; ok {
			return v
		}
		tMLastIndex := make(map[int]struct{})
		for li, _ := range mLastIndex {
			tMLastIndex[li-i] = struct{}{}
			tMLastIndex[li+i] = struct{}{}
			dp[li+i] = setMin(dp, li+i, i)
			dp[li-i] = setMin(dp, li-i, i)
		}
		mLastIndex = tMLastIndex
	}
}

func setMin(dp map[int]int, index, i int) int {
	if c, ok := dp[index]; ok {
		if c < i {
			return c
		}
	}
	return i
}
