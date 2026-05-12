type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "ぬ"
	}
	return strings.Join(strs, "あ")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "ぬ" {
		return []string{}
	}
	return strings.Split(encoded, "あ")
}
