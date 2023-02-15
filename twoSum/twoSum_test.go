package twoSum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		[]int{2, 7, 11, 15},
		[]int{3, 2, 4},
		[]int{3, 3},
	}
	targets := []int{
		9,
		6,
		6,
	}
	results := [][]int{
		[]int{0, 1},
		[]int{1, 2},
		[]int{0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := TowSum(tests[i], targets[i]); ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
