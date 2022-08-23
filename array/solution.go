package array

import "math/rand"

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	data := make([]int, len(nums))
	copy(data, nums)
	return Solution{data: nums}
}

func (s *Solution) Reset() []int {
	return s.data
}

func (s *Solution) Shuffle() []int {
	shuffle := make([]int, len(s.data))
	copy(shuffle, s.data)

	rand.Shuffle(len(shuffle), func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})
	return shuffle
}
