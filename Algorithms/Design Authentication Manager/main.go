package main

import "fmt"

func main() {
	// 0,0,0,2,2,2,3,3,3,3,3,3,3,4,5,5,5,6,7,7,7,7,7,7,7,8,8,8,8,9,9,8, 8,7,8,8,8,11,9,9,9,9,10,10,11,11,
	// 0,0,0,2,2,2,3,3,3,3,3,3,3,4,5,5,5,6,7,7,7,7,7,7,7,8,8,8,8,9,9,8, 7,6,7,7,7,9,7,7,7,7,8,8,9,9,
	a := Constructor(559)
	a.Renew("lv", 15)
	fmt.Print(a.CountUnexpiredTokens(26), ",")
	a.Renew("vttaj", 27)
	fmt.Print(a.CountUnexpiredTokens(36), ",")
	fmt.Print(a.CountUnexpiredTokens(37), ",")
	a.Generate("mib", 44)
	a.Generate("r", 55)
	fmt.Print(a.CountUnexpiredTokens(60), ",")
	fmt.Print(a.CountUnexpiredTokens(81), ",")
	fmt.Print(a.CountUnexpiredTokens(107), ",")
	a.Generate("jglr", 109)
	a.Renew("gewcr", 113)
	fmt.Print(a.CountUnexpiredTokens(140), ",")
	fmt.Print(a.CountUnexpiredTokens(143), ",")
	fmt.Print(a.CountUnexpiredTokens(162), ",")
	fmt.Print(a.CountUnexpiredTokens(181), ",")
	a.Renew("xwnxq", 192)
	fmt.Print(a.CountUnexpiredTokens(196), ",")
	fmt.Print(a.CountUnexpiredTokens(211), ",")
	a.Renew("pz", 236)
	a.Renew("k", 248)
	a.Renew("swptq", 249)
	fmt.Print(a.CountUnexpiredTokens(250), ",")
	a.Generate("ig", 288)
	fmt.Print(a.CountUnexpiredTokens(292), ",")
	a.Generate("iyzuk", 293)
	fmt.Print(a.CountUnexpiredTokens(305), ",")
	a.Renew("s", 312)
	fmt.Print(a.CountUnexpiredTokens(314), ",")
	fmt.Print(a.CountUnexpiredTokens(317), ",")
	a.Renew("jgs", 319)
	a.Generate("huyk", 321)
	fmt.Print(a.CountUnexpiredTokens(323), ",")
	a.Renew("kj", 348)
	a.Generate("rbww", 349)
	fmt.Print(a.CountUnexpiredTokens(371), ",")
	fmt.Print(a.CountUnexpiredTokens(372), ",")
	fmt.Print(a.CountUnexpiredTokens(387), ",")
	fmt.Print(a.CountUnexpiredTokens(409), ",")
	fmt.Print(a.CountUnexpiredTokens(416), ",")
	fmt.Print(a.CountUnexpiredTokens(440), ",")
	fmt.Print(a.CountUnexpiredTokens(456), ",")
	a.Generate("rvijr", 470)
	a.Renew("swptq", 478)
	a.Renew("iqx", 479)
	fmt.Print(a.CountUnexpiredTokens(484), ",")
	fmt.Print(a.CountUnexpiredTokens(489), ",")
	a.Renew("jgs", 496)
	fmt.Print(a.CountUnexpiredTokens(508), ",")
	fmt.Print(a.CountUnexpiredTokens(532), ",")
	a.Generate("dqjm", 553)
	fmt.Print(a.CountUnexpiredTokens(575), ",")
	fmt.Print(a.CountUnexpiredTokens(602), ",")
	fmt.Print(a.CountUnexpiredTokens(606), ",")
	a.Renew("wk", 625)
	a.Renew("dqjm", 647)
	fmt.Print(a.CountUnexpiredTokens(659), ",")
	fmt.Print(a.CountUnexpiredTokens(681), ",")
	a.Generate("wjj", 689)
	a.Renew("nsxwd", 706)
	a.Renew("qeqyy", 715)
	fmt.Print(a.CountUnexpiredTokens(728), ",")
	fmt.Print(a.CountUnexpiredTokens(759), ",")
	a.Renew("dfu", 773)
	fmt.Print(a.CountUnexpiredTokens(788), ",")
	a.Renew("vttaj", 807)
	a.Renew("dqjm", 816)
	a.Generate("gk", 823)
	a.Generate("mz", 838)
	fmt.Print(a.CountUnexpiredTokens(846), ",")
	fmt.Print(a.CountUnexpiredTokens(852), ",")
	fmt.Print(a.CountUnexpiredTokens(875), ",")
	a.Renew("pfu", 881)
	a.Generate("gkt", 895)
	fmt.Print(a.CountUnexpiredTokens(903), ",")
	a.Generate("pujx", 908)
	a.Renew("iflm", 917)
	fmt.Print(a.CountUnexpiredTokens(928), ",")
	a.Generate("k", 931)
	fmt.Print(a.CountUnexpiredTokens(942), ",")
	fmt.Print(a.CountUnexpiredTokens(948), ",")
	a.Renew("yx", 949)
	a.Generate("jywsu", 966)
	fmt.Print(a.CountUnexpiredTokens(977), ",")
	fmt.Print(a.CountUnexpiredTokens(985), ",")
	a.Generate("wk", 999)
}

type Node struct {
	TokenId     string
	ExpiredTime int
	Next        *Node
	Pre         *Node
}

type AuthenticationManager struct {
	TimeToLive int
	Len        int
	LastNode   *Node
	FirstNode  *Node
	NodesMap   map[string]*Node
}

func Constructor(timeToLive int) AuthenticationManager {
	firstNode, lastNode := &Node{}, &Node{}
	firstNode.Next = lastNode
	lastNode.Pre = firstNode
	return AuthenticationManager{
		TimeToLive: timeToLive,
		NodesMap:   make(map[string]*Node),
		FirstNode:  firstNode,
		LastNode:   lastNode,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.Clean(currentTime)

	node := &Node{
		TokenId:     tokenId,
		ExpiredTime: currentTime + this.TimeToLive,
	}
	prev := this.LastNode.Pre
	prev.Next = node
	node.Pre = prev
	node.Next = this.LastNode
	this.LastNode.Pre = node
	this.NodesMap[tokenId] = node
	this.Len++
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	this.Clean(currentTime)

	if node, ok := this.NodesMap[tokenId]; ok && node != nil && node.TokenId != "" {
		node.ExpiredTime = currentTime + this.TimeToLive
		// 把 node 移到链表最后
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		prev := this.LastNode.Pre
		prev.Next = node
		node.Pre = prev
		node.Next = this.LastNode
		this.LastNode.Pre = node
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	this.Clean(currentTime)
	// if currentTime == 659 {
	// 	for cur := this.FirstNode; cur.Next != nil; cur = cur.Next {
	// 		fmt.Println(currentTime, cur.TokenId, cur.ExpiredTime)
	// 	}
	// }
	return this.Len
}

func (this *AuthenticationManager) Clean(currentTime int) {
	if this.Len == 0 {
		return
	}
	for cur := this.FirstNode; cur.Next != nil; cur = cur.Next {
		if cur.TokenId == "" {
			continue
		}
		if cur.ExpiredTime > currentTime {
			break
		}
		if cur.ExpiredTime <= currentTime {
			this.NodesMap[cur.TokenId] = nil
			this.Len--
			cur.Pre.Next = cur.Next
			cur.Next.Pre = cur.Pre
			cur = cur.Pre
		}
	}
}
