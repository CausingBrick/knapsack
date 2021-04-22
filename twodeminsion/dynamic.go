package twodeminsion

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DynamicRec solution for the two-dimensional cost knapsack problem
func DynamicRec(items []*Item, kp *KnapSack) int {
	dp := make([][][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([][]int, kp.Capacity+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, kp.Volume+1)
		}
	}
	return dynamicRec(items, dp, kp.Capacity, kp.Volume, 0)
}

func dynamicRec(items []*Item, dp [][][]int, cap, vol, i int) int {
	value := 0
	if dp[i][cap][vol] > 0 {
		return value
	}

	switch {
	case i == len(items):
		break
	case items[i].Bulk > vol || items[i].Weight > cap:
		value = dynamicRec(items, dp, cap, vol, i+1)
	default:
		value = max(
			dynamicRec(items, dp, cap, vol, i+1),
			dynamicRec(
				items,
				dp,
				cap-items[i].Weight,
				vol-items[i].Bulk,
				i+1,
			)+items[i].Value,
		)
	}
	dp[i][cap][vol] = value
	return value
}

// Dynamic dynamic programing solution.
func Dynamic(items []*Item, kp *KnapSack) int {
	num := len(items)
	dp := make([][][]int, num+1)
	for i := range dp {
		dp[i] = make([][]int, kp.Capacity+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, kp.Volume+1)
		}
	}

	for i, item := range items {
		for j := 0; j < kp.Capacity; j++ {
			for k := 0; k < kp.Volume; k++ {
				if item.Bulk > kp.Volume || item.Weight > kp.Capacity {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = max(
						dp[i][j][k],
						dp[i][j-item.Weight][k-item.Bulk]+item.Value,
					)
				}
			}
		}
	}

	return dp[num][kp.Capacity][kp.Volume]
}

// DynamicCompress loop soulution with two-dimensional array.
func DynamicCompress(items []*Item, kp *KnapSack) int {
	dp := make([][]int, kp.Capacity+1)
	for i := range dp {
		dp[i] = make([]int, kp.Volume+1)
	}

	for _, item := range items {
		for j := kp.Capacity; j >= item.Weight; j-- {
			for k := kp.Volume; k >= item.Bulk; k-- {
				dp[j][k] = max(
					dp[j][k],
					dp[j-item.Weight][k-item.Bulk]+item.Value,
				)
			}
		}
	}

	return dp[kp.Capacity][kp.Volume]
}
