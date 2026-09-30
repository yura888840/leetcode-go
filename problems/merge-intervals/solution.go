package solution

import (
    "cmp"
    "slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	var (
		toBePushedToOutput []int
		output [][]int
	)
	intervals = sortIntervalsByStart(intervals)
	toBePushedToOutput = intervals[0]

	for _, currentInterval := range intervals[1:] {
		if intrvalsIntersects(toBePushedToOutput, currentInterval) {
			toBePushedToOutput = mergeIntervals(toBePushedToOutput, currentInterval)
		} else {
			output = append(output, toBePushedToOutput)
			toBePushedToOutput = currentInterval
		}
	}
	output = append(output, toBePushedToOutput)

	return output
}

func mergeIntervals(interval1, interval2 []int) []int {
	return []int{min(interval1[0], interval2[0]), max(interval1[1], interval2[1])}
}

func sortIntervalsByStart(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	return intervals
}

func intrvalsIntersects(interval1, interval2 []int) bool {
	s1 := interval1[0]
	e1 := interval1[1]

	s2 := interval2[0]
	e2 := interval2[1]

	return (s2 >= s1 && s2 <= e1) || ( s1 >= s2 && s1 <= e2)
}