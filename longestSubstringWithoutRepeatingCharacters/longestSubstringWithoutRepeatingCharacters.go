package longestSubstringWithoutRepeatingCharacters

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	//if s = "", maxLength = 0
	//if s = "a", ["a"]0
	//if s = "dcbac",
	//if s = "abba"
	var left, maxLength *int
	left = new(int)
	maxLength = new(int)
	//add and update hashMap as a cache
	hashMap := make(map[rune]int)
	for right, rune := range s {
		if _, exist := hashMap[rune]; exist {
			*left = right
		}
		hashMap[rune] = right
		if length := right - *left + 1; *maxLength < length {
			*maxLength = length
		}
	}
	return *maxLength
}
