func productExceptSelf(nums []int) []int {
    result := make([]int, 0, len(nums))
    
    prefix := 1
    for i := 0; i < len(nums); i++ {
        result = append(result, prefix)
        prefix *= nums[i]
    }

    suffix := 1
    for i := len(nums)-1; i >= 0; i-- {
        result[i] = result[i] * suffix
        suffix *= nums[i]
    }

    return result
}
