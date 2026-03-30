func hasDuplicate(nums []int) bool {

	numSet := make(map[int]bool)

	for _,value := range nums {
		_,ok := numSet[value]
		if ok {
			return true
		}
		numSet[value] = true
	}

	return false
    
}
