package hackerrank

import "testing"

func TestSimpleArraySum(t *testing.T) {
	res := simpleArraySum([]int{1, 2, 3})
	expected := 6
	if expected != res {
		t.Errorf("%d != %d", res, expected)
	}
}
