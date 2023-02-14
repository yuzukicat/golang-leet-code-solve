package main

import (
	"fmt"
)

func towSum(nums []int, target int) {
	numMap := make(map[int]int)
	solveSlice := []int{}
	for index, number := range nums {
		anotherNumber := target - number
		anotherIndex, exist := numMap[anotherNumber]
		numMap[number] = index
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
	nums := []int{3, 2, 4}
	target := 6
	towSum(nums, target)
}
