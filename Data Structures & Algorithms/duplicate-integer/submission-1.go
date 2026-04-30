func hasDuplicate(nums []int) bool {
    h := map[int]int{}

    for i, num := range nums {
        if _, ok := h[num]; ok {
            return true
        }

        h[num] = i
    }

    return false
}
