package main

import "fmt"

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{
		{"name", "bob"},
		{"age", "two"},
	}))
	fmt.Println(evaluate("hi(name)", [][]string{
		{"a", "b"},
	}))
	fmt.Println(evaluate("(a)(a)(a)aaa", [][]string{
		{"a", "yes"},
	}))
}

func evaluate(s string, knowledge [][]string) string {
	kl := make(map[string]string)
	for _, k := range knowledge {
		kl[k[0]] = k[1]
	}

	var (
		ans  []byte
		key  []byte
		flag bool
	)
	for _, b := range s {
		if b == ')' {
			if v, ok := kl[string(key)]; ok {
				ans = append(ans, []byte(v)...)
			} else {
				ans = append(ans, '?')
			}
			key = key[:0]
			flag = false
			continue
		}
		if b == '(' {
			flag = true
			continue
		}
		if flag {
			key = append(key, byte(b))
			continue
		}
		ans = append(ans, byte(b))
	}
	return string(ans)
}
