package main

import "fmt"

var (
	first = map[byte]struct{}{
		'q': struct{}{}, 'w': struct{}{}, 'e': struct{}{}, 'r': struct{}{}, 't': struct{}{},
		'y': struct{}{}, 'u': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'p': struct{}{},
		'Q': struct{}{}, 'W': struct{}{}, 'E': struct{}{}, 'R': struct{}{}, 'T': struct{}{},
		'Y': struct{}{}, 'U': struct{}{}, 'I': struct{}{}, 'O': struct{}{}, 'P': struct{}{},
	}
	second = map[byte]struct{}{
		'a': struct{}{}, 's': struct{}{}, 'd': struct{}{}, 'f': struct{}{}, 'g': struct{}{},
		'h': struct{}{}, 'j': struct{}{}, 'k': struct{}{}, 'l': struct{}{},
		'A': struct{}{}, 'S': struct{}{}, 'D': struct{}{}, 'F': struct{}{}, 'G': struct{}{},
		'H': struct{}{}, 'J': struct{}{}, 'K': struct{}{}, 'L': struct{}{},
	}
	third = map[byte]struct{}{
		'z': struct{}{}, 'x': struct{}{}, 'c': struct{}{}, 'v': struct{}{},
		'b': struct{}{}, 'n': struct{}{}, 'm': struct{}{},
		'Z': struct{}{}, 'X': struct{}{}, 'C': struct{}{}, 'V': struct{}{},
		'B': struct{}{}, 'N': struct{}{}, 'M': struct{}{},
	}
)

func findWords(words []string) []string {
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		flag := false
		alphabets := []byte(word)
		if _, ok := first[alphabets[0]]; ok {
			for _, alphabet := range alphabets {
				if _, ok := first[alphabet]; !ok {
					flag = true
					break
				}
			}
		}

		if _, ok := second[alphabets[0]]; ok {
			for _, alphabet := range alphabets {
				if _, ok := second[alphabet]; !ok {
					flag = true
					break
				}
			}
		}

		if _, ok := third[alphabets[0]]; ok {
			for _, alphabet := range alphabets {
				if _, ok := third[alphabet]; !ok {
					flag = true
					break
				}
			}
		}

		if !flag {
			result = append(result, word)
		}
	}
	return result
}

func main() {

	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
