func groupAnagrams(strs []string) [][]string {

    anaHashMap := make(map[[26]int][]string)

    for _,str := range strs {
        keyArr := [26]int{}
        for _,char := range str {
            keyArr[char - 'a']++ 
        }
        anaHashMap[keyArr] = append(anaHashMap[keyArr],str)
    }

    result := [][]string{}

    for _, listStrs := range anaHashMap {
    result = append(result, listStrs)
}
    return result
}
