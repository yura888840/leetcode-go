package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example_1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example_2_no_solution",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example_3_zeroes",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "single_triplet",
			nums: []int{3, -5, 2},
			want: [][]int{{-5, 2, 3}},
		},
		{
			name: "all_positive",
			nums: []int{1, 2, 3, 4},
			want: [][]int{},
		},
		{
			name: "all_negative",
			nums: []int{-5, -4, -3, -2, -1},
			want: [][]int{},
		},
		{
			name: "many_zeroes_no_duplicates",
			nums: []int{0, 0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "repeated_values_no_duplicates",
			nums: []int{-2, -2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "two_equal_values",
			nums: []int{-2, 1, 1},
			want: [][]int{{-2, 1, 1}},
		},
		{
			name: "cannot_reuse_same_index",
			nums: []int{-2, 1, 4},
			want: [][]int{},
		},
		{
			name: "multiple_solutions",
			nums: []int{-2, -1, 0, 1, 2},
			want: [][]int{{-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "solutions_without_zero",
			nums: []int{-4, -2, -1, 1, 2, 3},
			want: [][]int{
				{-4, 1, 3},
				{-2, -1, 3},
			},
		},
		{
			name: "mixed_duplicates",
			nums: []int{-2, 0, 0, 2, 2, -1, 1, 1},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
				{-1, 0, 1},
			},
		},
		{
			name: "large_values",
			nums: []int{-100000, 0, 100000, -100000, 100000},
			want: [][]int{{-100000, 0, 100000}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Решение может сортировать входной срез.
			input := append([]int(nil), tt.nums...)
			got := threeSum(input)

			for i, triplet := range got {
				if len(triplet) != 3 {
					t.Fatalf(
						"result[%d] = %v; want exactly 3 numbers",
						i, triplet,
					)
				}
			}

			if !reflect.DeepEqual(
				normalizeTriplets(got),
				normalizeTriplets(tt.want),
			) {
				t.Errorf(
					"threeSum(%v) = %v; want %v",
					tt.nums, got, tt.want,
				)
			}
		})
	}
}

// Сортирует копии, не изменяя результат функции.
// Не удаляет дубликаты: лишняя тройка должна провалить тест.
// nil и пустой срез считаются одинаковым пустым результатом.
func normalizeTriplets(triplets [][]int) [][]int {
	result := make([][]int, len(triplets))

	for i, triplet := range triplets {
		result[i] = append([]int(nil), triplet...)
		sort.Ints(result[i])
	}

	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})

	return result
}