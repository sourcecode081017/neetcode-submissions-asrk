func numUniqueEmails(emails []string) int {
    emailMap := make(map[string]bool)
    for _, email := range emails {
        emailMap[processEmailAddr(email)] = true
    }
    return len(emailMap)
}

func processEmailAddr(email string) string {
	splitStrings := strings.Split(email, "@")
    name := splitStrings[0]
    domain := splitStrings[1]
    var nameBuilder strings.Builder
    for _, c := range name {
        if c != '.' {
            nameBuilder.WriteRune(c)
        }
        if c == '+' {
            break;
        }
    }
    return nameBuilder.String() + domain
}