package leetcode

import (
	"reflect"
	"testing"
)

func TestSumRange(t *testing.T) {
	c := Constructor([]int{1, 2, 3})
	expected := 5
	res := c.SumRange(1, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRange0(t *testing.T) {
	c := Constructor([]int{1, 2, 3})
	expected := 6
	res := c.SumRange(0, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeSame(t *testing.T) {
	c := Constructor([]int{1, 2, 3})
	expected := 3
	res := c.SumRange(2, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
