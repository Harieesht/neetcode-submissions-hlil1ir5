func lengthOfLongestSubstring(s string) int {

	charHashSet := make(map[rune]bool)
	runed := []rune(s)
	resLen := 0
	left := 0

	for right := 0;right < len(s);right++ {

		for {
			_,ok := charHashSet[runed[right]]
			if ok {
				delete(charHashSet,runed[left])
				left++
			} else {
				charHashSet[runed[right]] = true
				resLen = max(resLen,(right-left+1))
				break
			}
		}

	}

	return resLen

}
