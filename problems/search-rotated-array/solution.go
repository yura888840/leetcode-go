package solution

func search(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return -1
	}

	pivot := searchPivot(nums)

	leftVirtual := 0
	rightVirtual := lenNums - 1 

	for leftVirtual <= rightVirtual {
		midVirtual := leftVirtual + (rightVirtual - leftVirtual) / 2
		midReal := (midVirtual + pivot) % lenNums

		if nums[midReal] < target {
			leftVirtual = midVirtual + 1
		} else if nums[midReal] > target {
			rightVirtual = midVirtual - 1
		} else if nums[midReal] == target {
			return midReal
		}
	}

	return -1
}

func searchPivot(nums []int) int{
	var (
		left int = 0
		right int = len(nums) - 1
	)

	for left < right {
		mid := left + (right - left) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}