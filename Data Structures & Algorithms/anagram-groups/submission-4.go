func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[[26]int][]string)
	resultArr := [][]string{}

	for _,str := range strs {
		var arr [26]int
		for _,c := range str {
			arr[c-'a']++
		}
		hashMap[arr] = append(hashMap[arr],str)
	}

	for _,value := range hashMap {
		resultArr = append(resultArr,value)
	}

	return resultArr
}
