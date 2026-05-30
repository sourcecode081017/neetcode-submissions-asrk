func isSubsequence(s string, t string) bool {
	var i, j int

	for i, j = 0, 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}
	return j == len(s)

}
