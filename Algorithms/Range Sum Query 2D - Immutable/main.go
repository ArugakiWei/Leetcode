package main

func main() {

}

type NumMatrix struct {
	Sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{
		Sum: make([][]int, len(matrix)+1),
	}
	nm.Sum[0] = make([]int, len(matrix[0])+1)
	for i, v := range matrix {
		nm.Sum[i+1] = make([]int, len(matrix[i])+1)
		for ii, vv := range v {
			nm.Sum[i+1][ii+1] = vv + nm.Sum[i+1][ii] + nm.Sum[i][ii+1] - nm.Sum[i][ii]
		}
	}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	x1, y1 := row1+1, col1+1
	x2, y2 := row2+1, col2+1
	return this.Sum[x2][y2] - this.Sum[x2][y1-1] - this.Sum[x1-1][y2] + this.Sum[x1-1][y1-1]
}
