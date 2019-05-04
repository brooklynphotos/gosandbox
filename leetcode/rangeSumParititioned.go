// https://leetcode.com/problems/range-sum-query-immutable/

package leetcode

type NumArrayParititoned struct {
	chunkSize int
	sums      [][]int
}

func PartitionedConstructor(nums []int) NumArrayParititoned {
	chunkCount := 4
	chunkSize := len(nums) / chunkCount
	partitionedSums := make([][]int, chunkCount)
	for c := 0; c < chunkCount; c++ {
		size := chunkSize
		if c+1 == chunkCount {
			size = len(nums) - (c * chunkSize)
		}
		sums := make([]int, size)
		start := chunkSize * c
		for i := range sums {
			n := nums[start+i]
			if i == 0 {
				sums[i] = n
			} else {
				sums[i] = sums[i-1] + n
			}
		}
		partitionedSums[c] = sums
	}
	return NumArrayParititoned{chunkSize, partitionedSums}
}

func (this *NumArrayParititoned) SumRange(i int, j int) int {
	return this.getValue(j) - this.getValue(i-1)
}

func (this *NumArrayParititoned) getValue(i int) int {
	if i < 0 {
		return 0
	}
	chunkId := i / this.chunkSize
	arrId := i % this.chunkSize
	// the last partition is the biggest
	maxSizePrevParititions := this.chunkSize * (len(this.sums) - 1)
	if i >= maxSizePrevParititions {
		chunkId = len(this.sums) - 1
		arrId = i - maxSizePrevParititions
	}
	return this.getMax(chunkId-1) + this.sums[chunkId][arrId]
}

/**
 * gets the max value up to the end of this chunk
 */
func (this *NumArrayParititoned) getMax(c int) int {
	if c < 0 {
		return 0
	}
	size := len(this.sums[c])
	return this.sums[c][size-1] + this.getMax(c-1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
