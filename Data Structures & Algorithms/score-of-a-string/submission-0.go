func scoreOfString(s string) int {
	sum := 0.0
	for i := 1; i < len(s); i = i + 1 {
		n1 := int(s[i])
		n2 := int(s[i-1])
		diff := n1 - n2
		sum += math.Abs(float64(diff))
	}
	return int(sum)
}
