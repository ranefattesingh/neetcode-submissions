type Solution struct{
    length int
}

func (s *Solution) Encode(strs []string) string {
   s.length = len(strs)
   var sb strings.Builder
    for _,  v := range strs {
        sb.WriteString(strconv.Itoa(len(v)))
        sb.WriteString("$")
        sb.WriteString(v)
        sb.WriteString("#")
    }
    
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    output := make([]string, 0, s.length)
    byteArr := make([]byte, 0, 3) // MAX 200 
    
    for i := 0; i < len(encoded); {
        if encoded[i] >= 48 && encoded[i] <= 57 {
            byteArr = append(byteArr, encoded[i])
            i++
        } else if encoded[i] == '#' {
            byteArr = make([]byte, 0, 3)
            i++
        } else {
            i++
            length, _ := strconv.Atoi(string(byteArr))
            output = append(output, encoded[i: i + length])
            i+= length
        }
    }
    
    return output
}