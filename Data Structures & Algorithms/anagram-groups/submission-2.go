func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[[26]int][]string)
    for _, v := range strs {
        key := [26]int{}
        for _, k := range v {
            key[k-'a']++
        }


        q, ok := hashMap[key]
        if !ok {
            q = make([]string, 0)
        }

        q = append(q, v)
        hashMap[key] = q
    }

    output := make([][]string, 0, len(hashMap))

    for _, v := range hashMap {
        output = append(output, v)
    }

    return output
}
