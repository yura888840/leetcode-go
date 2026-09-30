package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "example_1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "example_2_touching",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "example_3_unsorted",
			intervals: [][]int{{4, 7}, {1, 4}},
			want:      [][]int{{1, 7}},
		},
		{
			name:      "single_interval",
			intervals: [][]int{{2, 5}},
			want:      [][]int{{2, 5}},
		},
		{
			name:      "no_overlaps",
			intervals: [][]int{{1, 2}, {4, 5}, {7, 9}},
			want:      [][]int{{1, 2}, {4, 5}, {7, 9}},
		},
		{
			name:      "reverse_order",
			intervals: [][]int{{7, 9}, {4, 5}, {1, 2}},
			want:      [][]int{{1, 2}, {4, 5}, {7, 9}},
		},
		{
			name:      "nested_intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 8}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "same_start",
			intervals: [][]int{{1, 3}, {1, 8}, {1, 5}},
			want:      [][]int{{1, 8}},
		},
		{
			name:      "same_end",
			intervals: [][]int{{3, 8}, {1, 8}, {5, 8}},
			want:      [][]int{{1, 8}},
		},
		{
			name:      "duplicates",
			intervals: [][]int{{1, 4}, {1, 4}, {1, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "overlap_chain",
			intervals: [][]int{{1, 4}, {3, 6}, {5, 9}},
			want:      [][]int{{1, 9}},
		},
		{
			name:      "touching_chain",
			intervals: [][]int{{5, 7}, {1, 3}, {3, 5}},
			want:      [][]int{{1, 7}},
		},
		{
			name:      "bridge_between_groups",
			intervals: [][]int{{1, 3}, {6, 9}, {2, 7}},
			want:      [][]int{{1, 9}},
		},
		{
			name:      "multiple_merged_groups",
			intervals: [][]int{{10, 12}, {2, 6}, {11, 15}, {1, 3}, {20, 22}},
			want:      [][]int{{1, 6}, {10, 15}, {20, 22}},
		},
		{
			name:      "point_interval",
			intervals: [][]int{{3, 3}},
			want:      [][]int{{3, 3}},
		},
		{
			name:      "points_inside_and_on_boundaries",
			intervals: [][]int{{1, 1}, {1, 5}, {3, 3}, {5, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "separate_points",
			intervals: [][]int{{1, 1}, {2, 2}},
			want:      [][]int{{1, 1}, {2, 2}},
		},
		{
			name:      "gap_of_one_is_not_overlap",
			intervals: [][]int{{1, 4}, {5, 6}},
			want:      [][]int{{1, 4}, {5, 6}},
		},
		{
			name:      "zero_start",
			intervals: [][]int{{0, 0}, {0, 2}, {2, 4}},
			want:      [][]int{{0, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Глубокая копия: решение может менять входные интервалы.
			input := make([][]int, len(tt.intervals))
			for i, interval := range tt.intervals {
				input[i] = append([]int(nil), interval...)
			}

			got := merge(input)

			// Проверяем структуру до сортировки.
			for i, interval := range got {
				if len(interval) != 2 {
					t.Fatalf("result[%d] = %v; want [start, end]", i, interval)
				}
				if interval[0] > interval[1] {
					t.Fatalf("result[%d] has start > end: %v", i, interval)
				}
			}

			if !reflect.DeepEqual(
				sortedIntervals(got),
				sortedIntervals(tt.want),
			) {
				t.Errorf(
					"merge(%v) = %v; want %v",
					tt.intervals, got, tt.want,
				)
			}
		})
	}
}

// Сортирует копию внешнего среза, не меняя исходный результат.
func sortedIntervals(intervals [][]int) [][]int {
	result := append([][]int(nil), intervals...)

	sort.Slice(result, func(i, j int) bool {
		if result[i][0] == result[j][0] {
			return result[i][1] < result[j][1]
		}
		return result[i][0] < result[j][0]
	})

	return result
}