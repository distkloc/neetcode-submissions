func productExceptSelf(nums []int) []int {
	zeroCount := 0
	product := 1

	n := len(nums)
	res := make([]int, n)

    for _, v := range nums {
		if v == 0 {
			zeroCount++

			if zeroCount > 1 {
				return res
			}
			continue
		}

		product *= v
	}

	for i, v := range nums {
		if v == 0 {
			res[i] = product
			return res
		}

		if zeroCount == 0 {
			res[i] = product / v
		}
	}

	return res
}