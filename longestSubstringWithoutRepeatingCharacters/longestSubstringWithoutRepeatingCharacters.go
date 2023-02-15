package longestSubstringWithoutRepeatingCharacters

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	// s=""
	left := 0
	maxLength := 0
	//add and update hashMap as a cache
	hashMap := make(map[rune]int)
	for right, rune := range s {
		//Push the left to the latestPosition + 1
		if latestPosition, exist := hashMap[rune]; exist && latestPosition > left {
			left = latestPosition
		}
		if length := right - left + 1; maxLength < length {
			maxLength = length
		}
		//Optimize it here
		hashMap[rune] = right + 1
	}
	return maxLength
}
