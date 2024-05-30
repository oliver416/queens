package array

func CustomInversionSort(nums []int) []int {
    inversions := 0

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i+1] < nums[i] {
            nums[i], nums[i+1] = nums[i+1], nums[i]
            inversions++
        }

        if i == len(nums) - 2 {
            if inversions == 0 {
                break
            }

            inversions = 0
            i = 0 
        }
    }

    return nums
}

