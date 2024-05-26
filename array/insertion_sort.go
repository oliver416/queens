package array

func InsertionSort(nums []int) []int {
    for i, _ := range nums {
        for j := i; j > 0; j-- {
            if nums[j] < nums[j - 1] {
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
            }
        }
    }
    return nums
}

