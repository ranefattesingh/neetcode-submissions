func minWindow(s string, t string) string {
    window := make(map[byte]int)
    countT := make(map[byte]int)

    for _, i := range t {
        countT[byte(i)]++
        window[byte(i)] = 0
    }

    need := len(countT)
    have := 0
    left := 0
    minString := s
    found := false

    for right := 0; right < len(s); right++ {
        window[s[right]]++
        if countT[s[right]] > 0 && countT[s[right]] == window[s[right]] {
            have++
        }
        

        for have == need {
            found = true
            if len(s[left: right+1]) < len(minString) {
                minString = s[left: right+1]
            }

            window[s[left]]--
            if countT[s[left]] > 0 && window[s[left]] < countT[s[left]] {
                have--
            }

            left++
        }
    }

    if !found {
        return ""
    }

    return minString
}
