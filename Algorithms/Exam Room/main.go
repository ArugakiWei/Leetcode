package main

// not pass
func main() {
	e := Constructor(8)
	e.Seat()
	e.Seat()
	e.Seat()
}

type ExamRoom struct {
	sets []int
}

func Constructor(N int) ExamRoom {
	e := ExamRoom{
		sets: make([]int, N),
	}
	for i, _ := range e.sets {
		e.sets[i] = -1
	}

	return e
}

func (this *ExamRoom) Seat() int {
	start, end, maxLen, maxStart, maxEnd := 0, 0, -1, 0, 0
	for i := 0; i < len(this.sets); i++ {
		if this.sets[i] == -1 {
			start, end = i, i
			for {
				if end >= len(this.sets) || this.sets[end] != -1 {
					break
				}
				end++
			}
			i = end
			end--

			if this.sets[0] == -1 && this.sets[len(this.sets)-1] == -1 {
				if maxLen < end-start {
					maxLen = (end - start) / 2
					maxStart = start
					maxEnd = end
				}
			} else {
				if maxLen < (end-start)/2 {
					maxLen = (end - start) / 2
					maxStart = start
					maxEnd = end
				}
			}
		}
	}
	if maxStart == 0 && this.sets[0] == -1 {
		this.sets[0] = 1
		return 0
	}
	if maxEnd == len(this.sets)-1 && this.sets[len(this.sets)-1] == -1 {
		this.sets[len(this.sets)-1] = 1
		return len(this.sets) - 1
	}

	index := (maxEnd-maxStart)/2 + maxStart
	this.sets[index] = 1
	return index
}

func (this *ExamRoom) Leave(p int) {
	this.sets[p] = -1
}
