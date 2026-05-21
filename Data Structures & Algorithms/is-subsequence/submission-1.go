func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i = i + 1 {
		if t[i] == s[j] {
			j += 1
		}
	}
	return j == len(s)
}
