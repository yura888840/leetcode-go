package solution

import "testing"

func TestPlayground(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	t.Logf("Input:  %v", nums)

	got := threeSum(nums)

	t.Logf("Result: %v", got)
}