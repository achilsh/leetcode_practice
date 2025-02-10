package jianzhi_offer

import (
	"github.com/badgerodon/collections"
)

// QueueInStack 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead，分别完成在队列尾部插入节点和在队列头部删除节点的功能
// QueueInStack use two stack implement queue
// ctrl + . quick key for

type QueueImplementInStack struct {
	S1 collections.Stack[int]
	S2 collections.Stack[int]
}

func NewQueueImplementInStack() *QueueImplementInStack {
	item := new(QueueImplementInStack)
	item.S1 = collections.NewStack[int]()
	item.S2 = collections.NewStack[int]()
	return item
}

func (g *QueueImplementInStack) AppendTail(data int) {
	g.S1.Push(data)
}
// DeleteHead 删除数据只从第二个栈中删除，如果第二个没有数据则需要从第一个栈移栋数据。
func (g *QueueImplementInStack) DeleteHead() int {
	if g.S2.Size() == 0 {
		for {
			if g.S1.Size() != 0 {
				v, _ := g.S1.Pop()
				g.S2.Push(v)
			} else {
				break
			}
		}
	}
	if g.S2.Size() != 0 {
		v, _ := g.S2.Pop()
		return v
	}
	return -100

}
