func checkSpecial(c byte) bool {
	return c < '0' || c > 'Z' || (c > '9' && c < 'A')
}

func isPalindrome(s string) bool {
 	i := 0
 	j := len(s) - 1

 	uppered := strings.ToUpper(s)
 	for i <= j {
 		left := uppered[i]
 		right := uppered[j]

 		if checkSpecial(left) {
 			j++
 		} else if checkSpecial(right) {
 			i--
 		} else if left != right {
 			return false
 		}

 		i++
 		j--
 	}

    return true
}
