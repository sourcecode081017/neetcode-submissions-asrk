func validPalindrome(s string) bool {
	isPalindrome := func(str string, i, j int) bool {
		for i < j {
			if str[i] != str[j] {
				return false
			}

			i, j = i+1, j-1
		}
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
		}
	}
	return true

}
