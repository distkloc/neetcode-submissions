func maxProfit(prices []int) int {
	maxP := 0
	minPrice := prices[0]

	for _, p := range prices {
		maxP = max(maxP, p - minPrice)
		minPrice = min(minPrice, p)

		fmt.Println("p:", p, " maxP: ", maxP, " minPrice:", minPrice)
	}
	

	return maxP
}
