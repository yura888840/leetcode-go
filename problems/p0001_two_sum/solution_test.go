package solution

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{"basic", []int{2, 7, 11, 15}, 9},
		{"middle_and_end", []int{3, 2, 4}, 6},
		{"duplicates", []int{3, 3}, 6},
		{"negative_numbers", []int{-3, 4, 3, 90}, 0},
		{"zeroes", []int{0, 4, 3, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(append([]int(nil), tt.nums...), tt.target)
			if len(got) != 2 {
				t.Fatalf("twoSum(%v, %d) = %v; want two indices", tt.nums, tt.target, got)
			}
			i, j := got[0], got[1]
			if i < 0 || j < 0 || i >= len(tt.nums) || j >= len(tt.nums) || i == j {
				t.Fatalf("invalid indices: %v", got)
			}
			if tt.nums[i]+tt.nums[j] != tt.target {
				t.Errorf("nums[%d] + nums[%d] != %d", i, j, tt.target)
			}
		})
	}
}
