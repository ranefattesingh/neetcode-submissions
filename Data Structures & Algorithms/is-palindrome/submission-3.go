func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		a, valid := normalizeAlnum(s[left])
		if !valid {
			left++

			continue
		}

		b, valid := normalizeAlnum(s[right])
		if !valid {
			right--

			continue
		}

		if a != b {
			return false
		}

		left++
		right--
	}

	return true
}

func normalizeAlnum(c byte) (byte, bool) {
	if (c >= 'a' && c <= 'z') {
		return c, true
	}

	if (c >= 'A' && c <= 'Z')  {
		return c + ('a' - 'A'), true
	}
	
	if c >= '0' && c <= '9' {
		return c, true
	}

	return c, false
}

