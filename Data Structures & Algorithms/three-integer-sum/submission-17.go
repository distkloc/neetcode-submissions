func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }

		l := i + 1
		r := length - 1

		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]

			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
 
	return res
}
