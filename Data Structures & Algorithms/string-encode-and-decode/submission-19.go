type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := ""

	for _, s := range strs {
		res += strconv.Itoa(len(s)) + "#" + s
	}

	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}

	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length
	}
	return res
}
