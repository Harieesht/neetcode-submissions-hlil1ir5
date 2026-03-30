func topKFrequent(nums []int, k int) []int {

    numsLen := len(nums)

    countArr := make([][]int,numsLen+1)

    countMap := make(map[int]int)

    resultArr := []int{}

    for _,num := range nums {
        countMap[num]++
    }

    for key,value := range countMap {
        countArr[value] = append(countArr[value],key)
    }

    for i := len(countArr)-1 ; i>0 && len(resultArr)<k; i-- {
        arrValue := countArr[i]

        if len(arrValue) == 0 {
            continue
        } else {
            for _,value := range arrValue {
                if len(resultArr)< k {
                    resultArr = append(resultArr,value)
                }
            }
        }
    }

    return resultArr

}
