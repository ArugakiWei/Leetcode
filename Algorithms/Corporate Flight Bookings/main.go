package main

func main() {

}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := NewDiff(n)
	for _, v := range bookings {
		diff.Incr(v[0], v[1], v[2])
	}
	return diff.Result()
}

type Diff struct {
	Num []int
}

func NewDiff(n int) Diff {
	return Diff{
		Num: make([]int, n+1),
	}
}

func (d Diff) Incr(first, last, add int) {
	d.Num[first] += add
	if last+1 < len(d.Num) {
		d.Num[last+1] -= add
	}
}

func (d Diff) Result() []int {
	ans := make([]int, len(d.Num))
	ans[0] = d.Num[0]
	for i := 1; i < len(d.Num); i++ {
		ans[i] = d.Num[i] + ans[i-1]
	}
	return ans
}
