func isAnagram(s string, t string) bool {
    hashMap := make(map[rune]int)

    for _, v := range s {
        hashMap[v]++
    }

    for _, v := range t {
        hashMap[v]--
    }

    for _, v := range hashMap {
        if v != 0 {
            return false
        }
    }

    return true
}
