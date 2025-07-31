func palindrome(s string){
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s," ","")

	left := 0
	right := len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
		return true
	}
}