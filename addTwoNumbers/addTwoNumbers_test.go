package addTwoNumbers

import (
	"testing"
)

func generateNodeList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := new(ListNode)
	pointer := head
	for index, value := range list {
		pointer.Val = value
		if index != len(list)-1 {
			pointer.Next = new(ListNode)
			pointer = pointer.Next
		}
	}
	return head
}

func isEqual(left, right *ListNode) bool {
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	if left != nil || right != nil {
		return false
	}
	return true
}

func TestAddTwoNumbers(t *testing.T) {
	left := [][]int{
		[]int{2, 4, 3},
		[]int{0},
		[]int{9, 9, 9, 9, 9, 9, 9},
	}
	right := [][]int{
		[]int{5, 6, 4},
		[]int{0},
		[]int{9, 9, 9, 9},
	}
	results := [][]int{
		[]int{7, 0, 8},
		[]int{0},
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		l, r := generateNodeList(left[i]), generateNodeList(right[i])
		if ret := AddTwoNumbers(l, r); !isEqual(ret, generateNodeList(results[i])) {
			t.Fatalf("case %d fails: %v\n", i, ret.Val)
		}
	}
}
