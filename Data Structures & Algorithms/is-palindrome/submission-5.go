func isPalindrome(s string) bool {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    strippedS := re.ReplaceAllString(s, "")

    reverseS := ""

    for i := len(strippedS) - 1; i >= 0; i-- {
        reverseS += string(strippedS[i])
    }

    return strings.ToLower(reverseS) == strings.ToLower(strippedS)
}
