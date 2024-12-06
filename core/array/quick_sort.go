package array

func DivideArray(nums []int) ([]int, int) {
	pivot := nums[0]
	n := 0

	for i, number := range nums {
		if number < pivot {
			nums[i], nums[n] = nums[n], nums[i]
			n++
		}
	}

	return nums, n
}

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	nums, n := DivideArray(nums)

	QuickSort(nums[n+1:])
	QuickSort(nums[:n])

	return nums
}
