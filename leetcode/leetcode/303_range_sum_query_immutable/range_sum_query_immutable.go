package leetcode303

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{make([]int, n)}
	for i := 0; i < n; i++ {
		if i == 0 {
			na.sums[0] = nums[0]
		} else {
			na.sums[i] = na.sums[i-1] + nums[i]
		}
	}
	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
