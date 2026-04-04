func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }
    var res strings.Builder
    for i := 0; i < len(strs[0]); i++ {
        for _, s := range strs {
            if(i == len(s) || s[i] != strs[0][i]) {
                return res.String()
            }
        }
        res.WriteString(string(strs[0][i]))
    }
    return res.String()
}
