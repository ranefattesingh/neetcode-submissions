func twoSum(numbers []int, target int) []int {
    i := 0
    j := len(numbers) - 1

    for i < j {
        s := numbers[i] + numbers[j]
        if s == target {
            return []int{i+1, j+1}
        }

        if s < target {
            i++
        } else {
            j--
        }
    }

    return nil
}
