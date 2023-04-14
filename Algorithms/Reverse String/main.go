package main

func main() {
	reverseString([]byte("hello"))
}

func reverseString1(s string) string {

	var rs []byte
	for i := len(s) - 1; i >= 0; i-- {

		rs = append(rs, s[i])
	}
	return string(rs)
}

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
