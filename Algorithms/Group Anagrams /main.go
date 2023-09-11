package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"abbbbbbbbbbb", "aaaaaaaaaaab"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			count[122-c]++
		}
		k := toString(count)
		m[k] = append(m[k], s)
	}
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func toString(c []int) string {
	var b strings.Builder
	for _, n := range c {
		b.WriteString(strconv.Itoa(n))
		b.WriteString("-")
	}
	return b.String()
}
