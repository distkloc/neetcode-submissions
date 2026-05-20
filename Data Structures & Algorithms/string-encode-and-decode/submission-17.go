type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	sizes := make([]string, 0, len(strs))

	for _, s := range strs {
		sizes = append(sizes, strconv.Itoa(len(s)))
	}

	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return make([]string, 0, 0)
	}

	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")

	res := make([]string, 0, len(sizes))
	i := 0
	for _, size := range sizes {
		if size == "" {
			continue
		}

		length, _ := strconv.Atoi(size)
		res = append(res, parts[1][i:i+length])
		i += length
	}
	return res
}
