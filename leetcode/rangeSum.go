// https://leetcode.com/problems/range-sum-query-immutable/

package leetcode

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}
	return NumArray{nums}
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
