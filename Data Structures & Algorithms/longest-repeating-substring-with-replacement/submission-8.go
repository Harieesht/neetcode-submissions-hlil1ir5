func characterReplacement(s string, k int) int {
	freqCount := make(map[rune]int)
	runed := []rune(s)
	maxConsec := 0
	maxFreqCount := 0

	left := 0
	
	for right := 0;right < len(runed) ; right ++ {
		freqCount[runed[right]]++
		maxFreqCount = max(maxFreqCount, freqCount[runed[right]])
		
		conSec := right-left+1
		if conSec - maxFreqCount <= k {
			maxConsec = max(maxConsec,conSec)
		} else {
			freqCount[runed[left]]--
			left++
		}
	}

	return maxConsec
}