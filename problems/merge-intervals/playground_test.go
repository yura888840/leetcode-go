package solution

import (
	"fmt"
	"testing"
)

func TestPlayground(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	fmt.Println("Input:", intervals)

	result := merge(intervals)

	fmt.Println("Result:", result)
}