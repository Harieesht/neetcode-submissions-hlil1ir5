func maxProfit(prices []int) int {
	maxP := 0
	left := 0

	for index,value := range prices {
		if value < prices[left]{
			left = index
		} else {
			maxP = max(maxP,(prices[index]-prices[left]))
		}
	}

	return maxP

}
