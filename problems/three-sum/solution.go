package solution

import "slices"

func threeSum(nums []int) [][]int {
	seen := make(map[[3]int]bool)

	slices.Sort(nums[:])
	for num1_idx := 0; num1_idx < len(nums)-2; num1_idx++ {
		num1 := nums[num1_idx]

		left := num1_idx + 1
		right := len(nums) - 1

		for left < right {
			num2Candidate := nums[left]
			num3Candidate := nums[right]

			candidateSum := num1 + num2Candidate + num3Candidate
			if candidateSum == 0 {
				tripletFound := [3]int{num1, num2Candidate, num3Candidate}
				slices.Sort(tripletFound[:])

				if !seen[tripletFound] {
					seen[tripletFound] = true
				}
			}
			if candidateSum < 0 { 
				left++
			} else {
				right--
			}

		}
	}

	result := make([][]int, 0, len(seen))

	for triplet := range seen {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}

	return result
}

func _threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	counters := createCountersTable(nums)
	seen := make(map[[3]int]bool)

	for num1_idx := 0; num1_idx < len(nums)-2; num1_idx++ {
		num1 := nums[num1_idx]
		counters[num1]--

		for num2_idx := num1_idx + 1; num2_idx < len(nums); num2_idx++ {
			num2 := nums[num2_idx]

			counters[num2]--

			num3 := 0 - num1 - num2
			if counters[num3] > 0 {
				tripletFound := [3]int{num1, num2, num3}
				slices.Sort(tripletFound[:])

				if !seen[tripletFound] {
					seen[tripletFound] = true
				}
			}

			counters[num2]++
		}

		counters[num1]++
	}

	result := make([][]int, 0, len(seen))

	for triplet := range seen {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}

	return result
}

func createCountersTable(nums []int) map[int]uint {
	counters := make(map[int]uint)

	for _, num := range nums {
		counters[num]++
	}
	return counters
}
