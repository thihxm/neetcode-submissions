func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sFreq := make(map[string]int)
    tFreq := make(map[string]int)

    for i := 0; i < len(s); i++ {
        sFreq[string(s[i])]++
        tFreq[string(t[i])]++
    }

    for i := 0; i < len(s); i++ {
        if sFreq[string(s[i])] != tFreq[string(s[i])] {
            return false
        }        
    }

    return true
}
