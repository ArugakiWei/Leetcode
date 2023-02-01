package main

import "fmt"

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}

func decodeMessage(key string, message string) string {
	m := make(map[int32]int32)
	b := 'a'
	for _, v := range key {
		if v == ' ' {
			continue
		}
		if _, ok := m[v]; !ok {
			m[v] = b
			b++
		}
	}

	ans := make([]byte, 0, len(message))
	for _, v := range message {
		if v == ' ' {
			ans = append(ans, byte(v))
		} else {
			ans = append(ans, byte(m[v]))
		}
	}
	return string(ans)
}
