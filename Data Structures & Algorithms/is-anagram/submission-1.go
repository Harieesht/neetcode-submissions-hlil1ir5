func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    countMap := make(map[byte]int)

    for i:=0;i<len(s);i++ {
        countMap[s[i]]++
        countMap[t[i]]--
    }

    for _,value := range countMap {
        if value != 0{
            return false
        }
    }

    return true
    
}
