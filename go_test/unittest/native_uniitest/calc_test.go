package native_uniitest

import "testing"

func TestSum(t *testing.T) {
	rightResult := 8
	result := Sum(3, 5)
	if result != rightResult {
		t.Errorf("error, require %d, got %d", rightResult, result)
	}
}
