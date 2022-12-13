package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}

func checkIfPangram(sentence string) bool {
	m := map[byte]struct{}{
		'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}, 'h': {}, 'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {},
		'n': {}, 'o': {}, 'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {}, 'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
	}
	for _, i := range sentence {
		x := byte(i)
		if _, ok := m[x]; ok {
			delete(m, x)
		}
	}
	return len(m) == 0
}
