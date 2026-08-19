import "slices"

func isPalindrome(s string) bool {
	s = removeNonAlphanumeric(s)
	s = strings.ToLower(s)
	runes := []rune(s)

	slices.Reverse(runes)
	return string(runes) == s
}

func removeNonAlphanumeric(s string) string {
	return strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') {
				return r
			}
		return -1
	}, s)
}