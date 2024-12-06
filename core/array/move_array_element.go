package array

func MoveArrayElement(nums []int, from int, to int) []int {
	for i := from; i > to; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	return nums
}
