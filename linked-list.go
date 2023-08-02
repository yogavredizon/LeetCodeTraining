package leetcode

type any interface{}

type Node struct {
	Data any
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Add(list any) *Node {
	if l.Head == nil {
		l.Head.Data = list
	}

	if l.Head != nil {
		l.Head.Next.Data = list
	}

	return l.Head
}
