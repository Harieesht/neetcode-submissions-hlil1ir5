func twoSum(nums []int, target int) []int {
    
	numHash := make(map[int]int)

	for index,number := range nums {
		subValue := target-number
		value,ok := numHash[subValue]
		if ok {
			return []int{value,index}
		} 
		numHash[number] = index
	}

	return []int{-1,-1}
}
