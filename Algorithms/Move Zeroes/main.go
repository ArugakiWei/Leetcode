package main

func main() {

	a := []int{0, 1, 0, 3, 12}
	b := []int{0, 0, 1}
	c := []int{1, 0, 0}
	moveZeroes(a)
	moveZeroes(b)
	moveZeroes(c)
}

func moveZeroes1(nums []int) {

	j := 0
	for i := 0; j < len(nums); i++ {
		j++
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i--
		}
	}
}

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}
