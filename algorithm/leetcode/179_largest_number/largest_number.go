package leetcode179

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums) < 2 {
		return strconv.Itoa(nums[0])
	}
	strs := make([]string, len(nums))
	for index, item := range nums {
		strs[index] = strconv.Itoa(item)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
