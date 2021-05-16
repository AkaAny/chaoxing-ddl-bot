package utils

import "strconv"

func Max(nums ...int) int {
	var result = 0
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func ParseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
