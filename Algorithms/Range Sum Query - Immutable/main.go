package main

func main() {

}

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		Sum: make([]int, len(nums)+1),
	}
	for i, n := range nums {
		na.Sum[i+1] = na.Sum[i] + n
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Sum[right+1] - this.Sum[left]
}
