package array

func SelectionSort(nums []int) []int {
    for i, number := range(nums) {
        min_index := i
        min_number := number

        for j, sub := range(nums[i:]) {
            if sub < min_number {
                min_number = sub
                min_index = i + j
            }
        } 

        nums[i], nums[min_index] = nums[min_index], nums[i]
    }

    return nums
}

