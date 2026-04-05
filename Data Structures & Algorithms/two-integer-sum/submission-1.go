func twoSum(nums []int, target int) []int {
    prevMap := map[int]int{}

    for i, num := range nums {
        diff := target - num

        if val, ok := prevMap[diff]; ok {
            return []int{val, i}
        }
        prevMap[num] = i
    }

    return []int{}
}
