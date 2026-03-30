func twoSum(nums []int, target int) []int {

    numMap := make(map[int]int)

    for index,num := range nums {
        targetValue := target - num
        val,ok := numMap[targetValue]

        if !ok {
            numMap[num] = index
        } else {
            return []int{val,index}
        }
    }

    return []int{-1,-1}
    
}
