package solution

import "testing"

func TestSolve(t *testing.T) {
	// Replace the example cases and delete this Skip when starting.
	t.Skip("template: configure test cases and remove this line")

	tests := []struct {
		name string
		nums []int
		want int
	}{
		// Example shape only; adapt inputs and expected results to the problem.
		{"empty", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(append([]int(nil), tt.nums...))
			if got != tt.want {
				t.Errorf("solve(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
