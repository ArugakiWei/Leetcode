package main

import "fmt"

func main() {

	foo := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	bar := make(map[byte]string)
	for i, v := range foo{
		bar[97+byte(i)] = v
	}
	for i, k := range bar {
		fmt.Println(string(i), k)
	}
}


func uniqueMorseRepresentations(words []string) int {

}

