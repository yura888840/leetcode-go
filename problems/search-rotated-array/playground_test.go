package solution

import "testing"

func TestPlayground(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	t.Logf("Input: %v, target: %d", nums, target)

	got := search(nums, target)

	t.Logf("Result: %d", got)
}