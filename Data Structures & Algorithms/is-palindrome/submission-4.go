func isPalindrome(s string) bool {
    runes := make([]rune, 0, len(s))
    
    for _, c := range s {
        if unicode.IsDigit(c) || unicode.IsLetter(c) {
            runes = append(runes, unicode.ToLower(c))
        }
    }

    i, j := 0, len(runes)-1

    for i < j {
        if runes[i] != runes[j] {
            return false
        }
        i++
        j--
    }

    return true
}
