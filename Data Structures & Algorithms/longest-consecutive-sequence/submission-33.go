func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	longest := 0

	for _, v := range nums {
		numSet[v] = struct{}{}
	}

	for num := range numSet {
		if _, ok := numSet[num - 1]; !ok {
			length := 1

			for {
				if _, ok := numSet[num + length]; !ok {
					break
				}

				length++
			}

			longest = max(longest, length)
		}
	}

	return longest
}
