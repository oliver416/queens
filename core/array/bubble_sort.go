package array

func BubbleSort(nums []int) []int {
	for _ = range nums {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
