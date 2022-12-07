package main

import "fmt"

func main() {
	fmt.Println(numDifferentIntegers("a123bc34d8ef34"))
	fmt.Println(numDifferentIntegers("leet1234code234"))
	fmt.Println(numDifferentIntegers("a1b01c001"))
	fmt.Println(numDifferentIntegers("gi851a851q8510v"))
	fmt.Println(numDifferentIntegers("0a0"))
	fmt.Println(numDifferentIntegers("0i00e"))
}

func numDifferentIntegers(word string) int {
	m := make(map[string]struct{})
	var d []int32
	for _, b := range word {
		if b < 48 || b > 57 {
			if len(d) != 0 {
				m[num(d)] = struct{}{}
			}
			d = d[:0]
			continue
		}
		d = append(d, b)
	}
	if len(d) != 0 {
		m[num(d)] = struct{}{}
	}
	return len(m)
}

func num(b []int32) string {
	c := make([]int32, 0, len(b))
	var flag bool
	for _, v := range b {
		if flag {
			c = append(c, v)
			continue
		}
		if v != '0' {
			flag = true
			c = append(c, v)
		}
	}
	if !flag {
		return "0"
	}
	return string(c)
}
