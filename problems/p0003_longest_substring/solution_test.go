package solution

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// Delete this line when you start solving the problem.
	t.Skip("exercise: remove this line to enable tests")

	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"repeating_pattern", "abcabcbb", 3},
		{"all_same", "bbbbb", 1},
		{"substring_not_subsequence", "pwwkew", 3},
		{"do_not_move_left_backwards", "abba", 2},
		{"overlapping_window", "dvdf", 3},
		{"all_unique", "abcdef", 6},
		{"space", " ", 1},
		{"punctuation_and_spaces", "a! a", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}
