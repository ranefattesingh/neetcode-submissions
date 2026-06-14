type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := strings.Builder{}
	for _,  str := range strs {
		encoded.WriteString(strconv.Itoa(len(str)))
		encoded.WriteString("#")
		encoded.WriteString(str)
	}

	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	prev := 0
	decoded := make([]string, 0)
	for i := 0; i < len(encoded); i++ {
		if encoded[i] == '#' {
			num, _ := strconv.Atoi(encoded[prev:i])
			decoded = append(decoded, encoded[i+1: i + 1 + num])
			i = i + num
			prev = i + 1

		}
	}

	return decoded

}