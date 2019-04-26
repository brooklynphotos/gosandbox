package hackerrank

import "testing"

func TestPowerSum(t *testing.T) {
	res := powerSum(int32(100), int32(2), 1)
	expected := int32(3)
	if expected != res {
		t.Errorf("%d != %d", res, expected)
	}
}
