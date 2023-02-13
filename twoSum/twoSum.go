package main

import (
	"fmt"
)

func towSum(nums []int, target int) {
	numMap := make(map[int]int)
	for index, number := range nums {
		numMap[number] = index
	}
	for number, index := range numMap {
		anotherNumber := target - number
		anotherIndex, exist := numMap[anotherNumber]
		if exist && anotherIndex != index {
			fmt.Println(index, anotherIndex)
			break
		}
	}
}

func main() {
	nums := []int{2, 7, 11, 5}
	target := 9
	towSum(nums, target)
}
