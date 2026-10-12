package solution

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example_reachable",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "example_unreachable",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "single_zero_already_at_end",
			nums: []int{0},
			want: true,
		},
		{
			name: "single_positive_already_at_end",
			nums: []int{5},
			want: true,
		},
		{
			name: "cannot_leave_start",
			nums: []int{0, 1},
			want: false,
		},
		{
			name: "two_elements_reachable",
			nums: []int{1, 0},
			want: true,
		},
		{
			name: "all_ones",
			nums: []int{1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "all_zeroes",
			nums: []int{0, 0, 0},
			want: false,
		},
		{
			name: "zero_at_end_is_ok",
			nums: []int{1, 1, 1, 0},
			want: true,
		},
		{
			name: "zero_blocks_path",
			nums: []int{1, 0, 1},
			want: false,
		},
		{
			name: "jump_over_zero",
			nums: []int{2, 0, 0},
			want: true,
		},
		{
			name: "jump_over_multiple_zeroes",
			nums: []int{4, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "jump_falls_one_short",
			nums: []int{3, 0, 0, 0, 0},
			want: false,
		},
		{
			name: "maximum_jump_not_required",
			nums: []int{2, 3, 0, 0, 0},
			want: true,
		},
		{
			name: "jump_capacity_exceeds_array",
			nums: []int{10, 0, 0},
			want: true,
		},
		{
			name: "unreachable_large_jump",
			nums: []int{1, 0, 10, 0, 0},
			want: false,
		},
		{
			name: "several_jumps_over_zeroes",
			nums: []int{2, 0, 2, 0, 1, 0},
			want: true,
		},
		{
			name: "blocked_near_end",
			nums: []int{2, 0, 2, 0, 0, 1},
			want: false,
		},
		{
			name: "reachable_large_jump",
			nums: []int{1, 4, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "decreasing_jumps_reach_end",
			nums: []int{4, 3, 2, 1, 0},
			want: true,
		},
		{
			name: "decreasing_jumps_cannot_reach_end",
			nums: []int{4, 3, 2, 1, 0, 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)

			got := canJump(input)

			if got != tt.want {
				t.Errorf(
					"canJump(%v) = %t; want %t",
					tt.nums, got, tt.want,
				)
			}
		})
	}
}