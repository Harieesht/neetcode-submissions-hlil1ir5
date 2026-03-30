func productExceptSelf(nums []int) []int {

    numslen := len(nums)

    prefix := make([]int,numslen)

    suffix := make([]int,numslen)

    result := make([]int,numslen)

    for index,_ := range nums {
        if index == 0 {
            prefix[index] = nums[index]
            suffix[index+numslen-1] = nums[index+numslen-1]
            continue
        }
        prefix[index] = prefix[index-1] * nums[index]
        suffix[numslen-1-index] = nums[numslen-1-index] * suffix[numslen-index]
    }

    for index,_ := range result {
        if index == 0 {
            result[0] = suffix[1]
        } else if index == numslen - 1{
            result[index] = prefix[index-1]
        } else {
            result[index] = prefix[index-1] * suffix[index+1]
        }
    }

    return result



}
