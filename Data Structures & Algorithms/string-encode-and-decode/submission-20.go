type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s)))
		b.WriteString("#")
		b.WriteString(s)
	}

	return b.String()
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
