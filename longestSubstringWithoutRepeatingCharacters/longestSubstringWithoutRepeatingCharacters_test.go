package longestSubstringWithoutRepeatingCharacters

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"abba",
		"dvdf",
	}

	results := []int{
		3,
		1,
		3,
		2,
		3,
	}

	caseNum := 5

	for i := 0; i < caseNum; i++ {
		if ret := LongestSubstringWithoutRepeatingCharacters(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %d, expected: %d\n", i, ret, results[i])
		}
	}

}
