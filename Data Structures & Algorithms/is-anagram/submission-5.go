func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	charCount := make(map[byte]int)

	for i,_ := range s {
		charCount[s[i]] = charCount[s[i]] + 1
		charCount[t[i]] = charCount[t[i]] - 1
	}

	for _,value := range charCount {
		if value != 0 {
			return false
		}
	}

	return true
}
