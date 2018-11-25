package main

import (
	"fmt"
)

func main() {

	fmt.Println(isHappy(19))
}

//func isHappy(n int) bool {
//
//	runtime.GOMAXPROCS(4)
//
//	slowNum, fastNum := n, n
//	slowCh := make(chan int)
//	fastCh := make(chan int)
//	for  {
//		go slow(slowNum, slowCh)
//		go fast(fastNum, fastCh)
//		for i:=0; i<2; i++{
//			select {
//			case slowNum = <- slowCh :
//			case fastNum = <- fastCh :
//			}
//		}
//		if slowNum == fastNum {
//			break
//		}
//	}
//	if slowNum == 1 {
//		return true
//	}
//	return false
//}
//
//func slow(n int, ch chan int){
//	n = power(n)
//	ch <- n
//}
//
//func fast (n int, ch chan int){
//
//	n = power(n)
//	n = power(n)
//	ch <- n
//}
//
//func power(n int) int {
//	sum := 0
//	for {
//		i := n % 10
//		sum = sum + i*i
//		n = n / 10
//		if n == 0 {
//			break
//		}
//	}
//	return sum
//}



func isHappy(n int) bool {

	slow, fast := n, n
	for  {
		slow = power(slow)
		fast = power(fast)
		fast = power(fast)
		if slow == fast {
			break
		}
	}
	if slow == 1 {
		return true
	}
	return false
}

func power(n int) int {
	sum := 0
	for {
		i := n % 10
		sum = sum + i*i
		n = n / 10
		if n == 0 {
			break
		}
	}
	return sum
}