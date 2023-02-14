package main

import "fmt"

func iterate(input *string) {
	characterMap := make(map[rune]int)
	var leftIndex *int
	leftIndex = new(int)
	*leftIndex = 0
	var maxLength *int
	maxLength = new(int)
	*maxLength = 1
	for index, character := range *input {
		_, exist := characterMap[character]
		if exist {
			*leftIndex = index
		}
		windowLength := index - *leftIndex + 1
		if (windowLength) > *maxLength {
			*maxLength = windowLength
		}
		characterMap[character] = index
	}
	fmt.Println(*maxLength)
}

func main() {
	input1 := "abcabcbb"
	iterate(&input1)
	input2 := "bbbbb"
	iterate(&input2)
	input3 := "pwwkew"
	iterate(&input3)
	input4 := "aaaabcdefbcd"
	iterate(&input4)
	input5 := ""
	iterate(&input5)
}
