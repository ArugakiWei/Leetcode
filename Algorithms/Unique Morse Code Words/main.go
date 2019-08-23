package main

import "fmt"

func main() {

	b := []string{"zocd", "gjkl", "hzqk", "hzgq", "gjkl"}
	fmt.Println(uniqueMorseRepresentations(b))
}

func uniqueMorseRepresentations(words []string) int {

	foo := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	bar := make(map[byte]string)
	for i, v := range foo {
		bar[97+byte(i)] = v
	}
	mcode := make(map[string]interface{})
	for _, word := range words {
		cur := ""
		for _, w := range []byte(word) {
			cur = cur + bar[w]
		}
		mcode[cur] = new(interface{})
	}
	return len(mcode)
}
