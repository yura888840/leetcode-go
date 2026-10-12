package solution

// https://leetcode.com/problems/jump-game/description/
func canJump(nums []int) bool {
	maxReachableIndex := 0
	lenNums := len(nums)

	if lenNums == 1 {
		return true
	}
	for i := range nums {
		if i > maxReachableIndex {
			return false
		}

		if nums[i]+i > maxReachableIndex {
			maxReachableIndex = nums[i] + i
		}

		if maxReachableIndex >= (lenNums - 1) {
			return true
		}
	}

	return false
}
