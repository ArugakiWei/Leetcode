package main

import "fmt"

func main() {
	fmt.Println(equalFrequency("abcc"))
	fmt.Println(equalFrequency("aazz"))
	fmt.Println(equalFrequency("bbaz"))
	fmt.Println(equalFrequency("ddaccb"))
	fmt.Println(equalFrequency("abbcc"))
	fmt.Println(equalFrequency("zz"))
}

func equalFrequency(word string) bool {
	m := make(map[int32]int)
	for _, b := range word {
		m[b]++
	}
	mCount := make(map[int]int)
	for _, count := range m {
		mCount[count]++
	}
	if len(mCount) == 1 {
		if _, ok := mCount[1]; ok {
			return true
		}
		if _, ok := mCount[len(word)]; ok {
			return true
		}
	}
	if len(mCount) != 2 {
		return false
	}
	c1, ct1, c2, ct2 := -1, -1, -1, -1
	for count, countTime := range mCount {
		if c1 == -1 {
			c1, ct1 = count, countTime
		} else {
			c2, ct2 = count, countTime
		}
	}
	if (c1 == 1 && ct1 == 1) || (c2 == 1 && ct2 == 1) {
		return true
	}
	if c1-c2 == 1 && ct1 == 1 {
		return true
	}
	if c2-c1 == 1 && ct2 == 1 {
		return true
	}
	return false
}
