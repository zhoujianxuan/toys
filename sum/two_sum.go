package sum

func twoSum(nums []int, target int) []int {
	tMap := make(map[int]int)
	for i1, v := range nums {
		i2, ok := tMap[target-v]
		if ok {
			return []int{i1, i2}
		} else {
			tMap[v] = i1
		}
	}
	return []int{}
}
