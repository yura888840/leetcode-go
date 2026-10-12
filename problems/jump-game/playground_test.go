package solution

import "testing"

func TestPlayground(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}

	t.Logf("Input: %v", nums)

	got := canJump(nums)

	t.Logf("Result: %t", got)
}