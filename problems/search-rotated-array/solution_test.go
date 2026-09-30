package solution

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example_found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "example_not_found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "single_not_found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "single_found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "two_sorted_first",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "two_sorted_last",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "two_rotated_first",
			nums:   []int{3, 1},
			target: 3,
			want:   0,
		},
		{
			name:   "two_rotated_last",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "two_rotated_missing",
			nums:   []int{3, 1},
			target: 2,
			want:   -1,
		},
		{
			name:   "not_rotated_middle",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 4,
			want:   3,
		},
		{
			name:   "not_rotated_first",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 0,
			want:   0,
		},
		{
			name:   "not_rotated_last",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 7,
			want:   6,
		},
		{
			name:   "not_rotated_missing_inside",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 3,
			want:   -1,
		},
		{
			name:   "rotated_first",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			want:   0,
		},
		{
			name:   "rotated_last",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
		{
			name:   "maximum_before_rotation_boundary",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			want:   3,
		},
		{
			name:   "inside_left_part",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			want:   1,
		},
		{
			name:   "inside_right_part",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 1,
			want:   5,
		},
		{
			name:   "rotation_by_one",
			nums:   []int{1, 2, 3, 4, 5, 6, 0},
			target: 0,
			want:   6,
		},
		{
			name:   "rotation_by_n_minus_one",
			nums:   []int{6, 0, 1, 2, 3, 4, 5},
			target: 6,
			want:   0,
		},
		{
			name:   "small_left_part_search_right",
			nums:   []int{6, 0, 1, 2, 3, 4, 5},
			target: 3,
			want:   4,
		},
		{
			name:   "even_length",
			nums:   []int{4, 5, 0, 1, 2, 3},
			target: 0,
			want:   2,
		},
		{
			name:   "negative_target",
			nums:   []int{0, 3, 8, -10, -5, -2},
			target: -5,
			want:   4,
		},
		{
			name:   "below_minimum",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: -1,
			want:   -1,
		},
		{
			name:   "above_maximum",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 8,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)

			got := search(input, tt.target)

			if got != tt.want {
				t.Errorf(
					"search(%v, %d) = %d; want %d",
					tt.nums, tt.target, got, tt.want,
				)
			}
		})
	}
}