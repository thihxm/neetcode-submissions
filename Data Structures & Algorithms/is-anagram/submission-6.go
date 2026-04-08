func isAnagram(s string, t string) bool {
    if len(t) != len(s) {
        return false
    }

    freq := make(map[rune]int)

    for _, char := range s {
        freq[char]++
    }
    
    for _, char := range t {
        freq[char]--
        if (freq[char] < 0) {
            return false
        }
    }

    return true
}
