package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	memo := make(map[string]bool)
	fill(image, sr, sc, oldColor, newColor, memo)
	return image
}

func fill(image [][]int, sr int, sc int, oldColor, newColor int, memo map[string]bool) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[sr]) {
		return
	}
	if image[sr][sc] != oldColor {
		return
	}
	if v, ok := memo[fmt.Sprintf("%d%d", sr, sc)]; ok && v {
		return
	}

	image[sr][sc] = newColor
	memo[fmt.Sprintf("%d%d", sr, sc)] = true
	fill(image, sr+1, sc, oldColor, newColor, memo)
	fill(image, sr-1, sc, oldColor, newColor, memo)
	fill(image, sr, sc+1, oldColor, newColor, memo)
	fill(image, sr, sc-1, oldColor, newColor, memo)
}
