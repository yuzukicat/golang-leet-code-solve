package twoSum

func TowSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, number := range nums {
		if anotherIndex, exist := hashMap[number]; exist {
			return []int{anotherIndex, index}
		}
		//Store opposite in hashMapcache
		hashMap[target-number] = index
	}
	return nil
}
