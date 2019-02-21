package leetcode001

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	nums := []int{1, 3, 4, 5, 7, 10}
	result := twoSum(nums, 10)
	resultBs := twoSumBestSolution(nums, 10)
	if !reflect.DeepEqual(result, []int{1, 4}) {
		t.Error("twoSUm Error")
	}
	if !reflect.DeepEqual(resultBs, []int{1, 4}) {
		t.Error("twoSumBestSolution Error")
	}

}
