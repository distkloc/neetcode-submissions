import "slices"

func longestConsecutive(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}

	slices.Sort(nums)
	res := 1
	streak := 1
	fmt.Printf("%v", nums)
	for i := 0; i < length - 1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}

		if nums[i] + 1 != nums[i+1] {
			streak = 1
			continue
		}

		streak++

		if streak > res {
			res = streak
		}
	}
	return res
}
