package leetcode_test

import (
	"testing"

	lc "github.com/yogavredizon/LeetCodeTraining"
)

func TestMiddleNode(t *testing.T) {
	type LinkNode struct {
		next []int
	}
	testTable := []struct {
		name            string
		value           *LinkNode
		expectedOutCome []int
	}{
		{name: "First Chain", value: &LinkNode{next: []int{1, 2, 3, 4, 5}}, expectedOutCome: []int{4, 5, 6}},
		{name: "First Chain", value: &LinkNode{next: []int{1, 2, 3, 4}}, expectedOutCome: []int{3, 4}},
		{name: "First Chain", value: &LinkNode{next: []int{1, 2, 3}}, expectedOutCome: []int{2, 3}},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			want := lc.MiddleNode(&lc.ListNode{})
		})
	}
}
