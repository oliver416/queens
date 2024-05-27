package array

func Merge(arr1 []int, arr2 []int) []int {
    result := []int{}
    i, j := 0, 0

    for k := 0; k < len(arr1) + len(arr2); k++ {
        if j == len(arr2) || (i < len(arr1) && arr1[i] < arr2[j]) {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    } 

    return result
}

func MergeSort(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }

    if len(nums) == 2 {
        return Merge(nums[:1], nums[1:])
    }

    half := len(nums) / 2

    return Merge(
        MergeSort(nums[:half]),
        MergeSort(nums[half:]),
    )
}

