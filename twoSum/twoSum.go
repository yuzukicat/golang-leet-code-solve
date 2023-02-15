package twoSum

func TowSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, number := range nums {
		_, exist := hashMap[number]
		if exist {
			return []int{hashMap[number], index}
		}
		hashMap[target-number] = index
	}
	return nil
}
