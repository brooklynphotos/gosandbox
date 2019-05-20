// https://www.hackerrank.com/challenges/big-sorting/problem
package hackerrank

import "sort"

type byBigNumber []string

func (arr byBigNumber) Len() int {
	return len(arr)
}

func (arr byBigNumber) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr byBigNumber) Less(i, j int) bool {
	x := arr[i]
	y := arr[j]
	if len(x) < len(y) {
		return true
	}
	if len(x) > len(y) {
		return false
	}
	yrunes := []rune(y)
	for ind, xv := range x {
		yv := yrunes[ind]
		if xv < yv {
			return true
		}
		if xv > yv {
			return false
		}
	}
	return false
}

func bigSorting(unsorted []string) []string {
	sort.Sort(byBigNumber(unsorted))
	return unsorted

}
