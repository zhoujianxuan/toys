package no3

func GetNo() int32 {
	return 3
}

// 滑动窗口的方式遍历字符串，[left,right]保证内部字符串都不相同时为最大连着的字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var left int
	var right int
	var max int
	for right <= len(s) {
		if left == right {
			right++
			continue
		}
		ok := verify(s[left:right])
		if ok {
			if max < right-left {
				max = right - left
			}
			right++
		} else {
			left++
		}
	}
	return max
}

func verify(subString string) bool {
	if len(subString) == 0 {
		return true
	}
	for i, s := range subString {
		for j := int(i) + 1; j < len(subString); j++ {
			s2 := subString[j]
			if s == int32(s2) {
				return false
			}
		}
	}
	return true
}
