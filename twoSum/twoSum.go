package main

import (
	"fmt"
)

func towSum(nums []int, target int) {
	numMap := make(map[int]int)
	solveSlice := []int{}
	for index, number := range nums {
		anotherNumber := target - number
		numMap[number] = index
		anotherIndex, exist := numMap[anotherNumber]
		if exist && anotherIndex != index {
			solveSlice = append(solveSlice, anotherIndex)
			solveSlice = append(solveSlice, index)
			break
		} else if exist && anotherIndex == index {
			solveSlice = append(solveSlice, index)
		}
	}
	fmt.Println(solveSlice)
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	towSum(nums, target)
}
