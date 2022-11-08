package main

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	// fmt.Println(maximumSubarraySum([]int{1, 1, 1, 7, 8, 9}, 3))
}

func maximumSubarraySum(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}
