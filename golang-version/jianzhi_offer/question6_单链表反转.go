package jianzhi_offer

import "fmt"

type SingleLink struct {
	Next *SingleLink
	Data int
}

// NewSingleLinkBySlice 根据分片值创建一个单向链表
func NewSingleLinkBySlice(data []int) *SingleLink {
	if len(data) <= 0 {
		return nil
	}
	var head *SingleLink = nil
	var iterNode *SingleLink
	for i := 0; i < len(data); i++ {
		tmpNode := &SingleLink{
			Next: nil,
			Data: data[i],
		}
		fmt.Printf("creat node: %p, data: %v\n", tmpNode, data[i])

		if head == nil {
			head, iterNode = tmpNode, tmpNode
			continue
		}
		iterNode.Next = tmpNode
		iterNode = tmpNode
	}
	return head
}

func IterSingleLink(head *SingleLink) []int {
	var ret []int
	for {
		if head != nil {
			fmt.Printf("data: %v, current addr: %p\n", head.Data, head)

			ret = append(ret, head.Data)
			head = head.Next
			continue
		}
		break
	}
	return ret
}

// ReverseSingleLink 反转链表
// p2 = p1.Next; p3= p2.Next
// p1 = p2.Next;
func ReverseSingleLink(head *SingleLink) *SingleLink {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := p1.Next
	p3 := p2.Next
	for {
		p2.Next = p1
		if head.Next != nil {
			head.Next = nil
		}

		p1 = p2
		p2 = p3
		if p2 == nil {
			return p1
		}
		p3 = p2.Next
	}
}
