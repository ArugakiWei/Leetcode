package main

import "strings"

func main() {

}

type StreamChecker struct {
	Words  []string
	Stream strings.Builder
}

func Constructor(words []string) StreamChecker {
	return StreamChecker{
		Words: words,
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.Stream.WriteByte(letter)
	for _, word := range this.Words {
		if strings.HasSuffix(this.Stream.String(), word) {
			return true
		}
	}
	return false
}
