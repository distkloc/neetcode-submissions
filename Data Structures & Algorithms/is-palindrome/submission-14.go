func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		

		for l < r && !isAlphanumeric(rune(s[l])) {
			l++
		}

		for r > l && !isAlphanumeric(rune(s[r])) {
			r--
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		l++
		r--
	}
	
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}