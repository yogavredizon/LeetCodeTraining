package leetcode

type ListNode struct {
	Next []int
}

func MiddleNode(head *ListNode) *ListNode {
	store := make([]*ListNode, 10)

	if head != nil {
		store = append(store, head.Next)
	}

	return store[len(store)/2]
}
