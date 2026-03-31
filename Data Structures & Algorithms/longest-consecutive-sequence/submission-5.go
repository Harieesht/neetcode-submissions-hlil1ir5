func longestConsecutive(nums []int) int {

    numSet := make(map[int]bool)
    maxConSec := 0

    for _,num := range nums {
        numSet[num] = true
    }

    for key,_ := range numSet {

        conSec := 1
        curr := key

        if _, ok := numSet[key-1]; ok {
            continue
        }


        for {
        _,ok := numSet[curr+1]
        if !ok {
            break
        } else {
            conSec++
            curr++
        }}
        maxConSec = max(maxConSec,conSec)
    }

    return maxConSec


}
