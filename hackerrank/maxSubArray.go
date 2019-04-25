package hackerrank

import (
	"math"
	"sort"
)

func maxSubarray(arr []int32) []int32 {
	return []int32{
		maxSarray(arr),
		maxSubsequence(arr),
	}
}

func maxSarray(arr []int32) int32 {
	min := math.MinInt32
	maxCh := make(chan int32, len(arr))
	for i := 0; i < len(arr); i++ {
		go func(start int) {
			maxSum := int32(min)
			curSum := int32(0)
			for j := start; j < len(arr); j++ {
				curSum += arr[j]
				maxSum = max(maxSum, curSum)
			}
			maxCh <- maxSum
		}(i)
	}
	maxSum := int32(min)
	for i := 0; i < len(arr); i++ {
		maxSum = max(maxSum, <-maxCh)
	}
	return maxSum
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func buildMatrix(size int) [][]int32 {
	dp := make([][]int32, size)
	for i := range dp {
		dp[i] = make([]int32, size)
	}
	return dp
}

func maxSubsequence(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	if lastEl := arr[len(arr)-1]; lastEl <= 0 {
		return lastEl
	}
	sum := int32(0)
	for _, x := range arr {
		if sum > 0 || x > 0 {
			sum += x
		}
	}
	return sum
}
