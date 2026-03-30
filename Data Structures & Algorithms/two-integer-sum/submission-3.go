func twoSum(nums []int, target int) []int {
    
    numMap := make(map[int]int)

    for numIndex,num := range nums {
        subVal := target - num
        if index,exists := numMap[subVal]; exists {
            result := []int{index,numIndex}
            return result
        }
        numMap[num] = numIndex
    }

    return []int{0,0}

}
