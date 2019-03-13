package main

import (
	"fmt"
	"strings"
)

type LRUCache struct {

	cap		int
	data	map[int]*Node
	list	List
}

type Node struct {

	key 	int
	value 	int

	pre 	*Node
	next	*Node
}

func NewNode(key, value int) *Node {

	return &Node{
		key: 	key,
		value: 	value,
	}
}

type List struct {

	start	*Node
	end		*Node
}

func NewList() List {

	start := NewNode(-1, -1)
	end	  := NewNode(-2, -2)

	start.pre = nil
	end.next = nil
	start.next = end
	end.pre = start

	return List{
		start:	start,
		end:	end,
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

func (l List) movetoStart(node *Node) {

	temp := NewNode(node.key, node.value)
	l.puttoStart(temp)

	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil
	node.next = nil
	node = nil
}


func Constructor(capacity int) LRUCache {

	return LRUCache{
		cap:		capacity,
		data:		make(map[int]*Node),
		list:		NewList(),
	}
}

func (this *LRUCache) Get(key int) int {

	if node, ok := this.data[key]; ok {
		this.list.movetoStart(node)
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {

	if node, ok := this.data[key]; ok {
		node.value = value
		this.data[key] = node
		this.list.movetoStart(node)
	}else{
		if len(this.data) + 1 > this.cap {
			node := this.list.pop()
			delete(this.data, node.key)
		}
		node := NewNode(key, value)
		this.data[key] = node
		this.list.puttoStart(node)
	}
}

// cap == 1
func main() {

	cache := Constructor(5)

	cache.Put(1, 1);
	fmtlist(cache.list)

	cache.Put(2, 2);
	fmtlist(cache.list)

	cache.Put(3, 3);
	fmtlist(cache.list)

	cache.Put(4, 4);
	fmtlist(cache.list)

	fmt.Println(cache.Get(1))
	fmtlist(cache.list)

	fmt.Println(cache.Get(3))
	fmtlist(cache.list)

	fmt.Println(cache.Get(4))
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
			fmt.Printf("{%d:%d}" , node.key, node.value)
			break
		}
		fmt.Printf("{%d:%d}<-->" , node.key, node.value)
		node = node.next
	}
	fmt.Println()
}

func foo() ([]string, []string){
	a := `["put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]`
	b := `[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]`

	a = strings.Replace(a, `"`, "", -1)
	a = strings.Replace(a, `[`, "", -1)
	a = strings.Replace(a, `]`, "", -1)
	aa := strings.Split(a, ",")

	bb := strings.Split(b, "],")
	var ss []string

	for _, bbb := range bb {
		bbb = strings.Replace(bbb, `]`, "", -1)
		bbb = strings.Replace(bbb, `[`, "", -1)
		ss = append(ss, bbb)
	}


	return aa, ss
}