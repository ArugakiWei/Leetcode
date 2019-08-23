package main

import (
	"fmt"
)

type LRUCache struct {
	cap  int
	data map[int]*Node
	list List
}

type Node struct {
	key   int
	value int

	pre  *Node
	next *Node
}

func NewNode(key, value int) *Node {

	return &Node{
		key:   key,
		value: value,
	}
}

type List struct {
	start *Node
	end   *Node
}

func NewList() List {

	start := NewNode(-1, -1)
	end := NewNode(-2, -2)

	start.pre = nil
	end.next = nil
	start.next = end
	end.pre = start

	return List{
		start: start,
		end:   end,
	}
}

func (l List) puttoStart(node *Node) {

	node.pre = l.start
	node.next = l.start.next
	l.start.next = node
	node.next.pre = node
}

func (l List) pop() *Node {

	node := l.end.pre
	l.end.pre.pre.next = l.end
	l.end.pre = l.end.pre.pre
	return node
}

func (l List) movetoStart(node *Node) *Node {

	temp := NewNode(node.key, node.value)
	l.puttoStart(temp)

	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil
	node.next = nil
	node = nil

	return temp
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		cap:  capacity,
		data: make(map[int]*Node),
		list: NewList(),
	}
}

func (this *LRUCache) Get(key int) int {

	if node, ok := this.data[key]; ok {
		newNode := this.list.movetoStart(node)
		this.data[key] = newNode
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.data[key]; ok {
		node.value = value
		newNode := this.list.movetoStart(node)
		this.data[key] = newNode
	} else {
		if len(this.data)+1 > this.cap {
			node := this.list.pop()
			delete(this.data, node.key)
		}
		node := NewNode(key, value)
		this.list.puttoStart(node)
		this.data[key] = node
	}
}

// cap == 1
func main() {

	cache := Constructor(5)

	cache.Put(1, 1)
	fmtlist(cache.list)

	fmt.Println(cache.Get(1))
	fmtlist(cache.list)

	cache.Put(1, 3)
	fmtlist(cache.list)

	fmt.Println(cache.Get(1))
	fmtlist(cache.list)

	cache.Put(1, 4)
	fmtlist(cache.list)

	fmt.Println(cache.Get(1))
	fmtlist(cache.list)

	cache.Put(1, 5)
	fmtlist(cache.list)

	fmt.Println(cache.Get(1))
	fmtlist(cache.list)
	//
	//cache.Put(4, 4)
	//fmtlist(cache.list)
	//
	//fmt.Println(cache.Get(1))
	//fmtlist(cache.list)
	//
	//fmt.Println(cache.Get(3))
	//fmtlist(cache.list)
	//
	//fmt.Println(cache.Get(4))
	//fmtlist(cache.list)
	//ops, para := foo()
	//
	//for index, op := range ops {
	//	if op == "put" {
	//		kv := strings.Split(para[index], ",")
	//		k, _ := strconv.Atoi(kv[0])
	//		v, _ := strconv.Atoi(kv[1])
	//		cache.Put(k, v)
	//		fmtlist(cache.list)
	//	}
	//	if op == "get" {
	//		k, _ := strconv.Atoi(para[index])
	//		cache.Get(k)
	//		fmtlist(cache.list)
	//	}
	//	time.Sleep(2 * time.Second)
	//}
}

func fmtlist(list List) {
	node := list.start
	for {
		if node.next == nil {
			fmt.Printf("{%d:%d}", node.key, node.value)
			break
		}
		fmt.Printf("{%d:%d}<-->", node.key, node.value)
		node = node.next
	}
	fmt.Println()
}
