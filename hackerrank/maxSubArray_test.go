package hackerrank

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	res := maxSubarray([]int32{1, 2, 3, 4})
	expected := []int32{10, 10}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
