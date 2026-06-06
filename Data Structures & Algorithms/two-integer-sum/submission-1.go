func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        j, ok := hashMap[nums[i]]
        if ok {
            return []int{j, i}
        }

        hashMap[target-nums[i]] = i
    }

    return nil
}
