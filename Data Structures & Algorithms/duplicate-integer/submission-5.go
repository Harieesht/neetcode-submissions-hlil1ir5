func hasDuplicate(nums []int) bool {
    
	numSet := make(map[int]struct{})

	for i := 0;i<len(nums);i++ {
		if _,exists := numSet[nums[i]]; exists {
			return true
		}
		
		numSet[nums[i]] = struct{}{};
		
	}

	return false
}
