package leetcode

import (
	"reflect"
	"testing"
)

func TestSumRangeParitioned(t *testing.T) {
	c := PartitionedConstructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := 5
	res := c.SumRange(1, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitioned0(t *testing.T) {
	c := PartitionedConstructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := 6
	res := c.SumRange(0, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitionedSame(t *testing.T) {
	c := PartitionedConstructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := 3
	res := c.SumRange(2, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitionedPerfect(t *testing.T) {
	c := PartitionedConstructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	expected := 78
	res := c.SumRange(0, 11)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitionedLC1(t *testing.T) {
	c := PartitionedConstructor([]int{-2, 0, 3, -5, 2, -1})
	expected := 1
	res := c.SumRange(0, 2)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitionedLC2(t *testing.T) {
	c := PartitionedConstructor([]int{-2, 0, 3, -5, 2, -1})
	expected := -1
	res := c.SumRange(2, 5)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
func TestSumRangeParitionedLC3(t *testing.T) {
	c := PartitionedConstructor([]int{-2, 0, 3, -5, 2, -1})
	expected := -3
	res := c.SumRange(0, 5)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%d != %d", res, expected)
	}
}
