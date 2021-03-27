package knapsack01

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DynamicRecursive dynamic programing recursive solution
func DynamicRecursive(profits []int, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	return dynamicRecursive(profits, weights, dp, 0, capacity)
}

func dynamicRecursive(pft, wgt []int, dp [][]int, i, cap int) int {
	if dp[i][cap] > 0 {
		return dp[i][cap]
	}
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = dynamicRecursive(pft, wgt, dp, i+1, cap)
	} else {
		value = max(dynamicRecursive(pft, wgt, dp, i+1, cap),
			dynamicRecursive(pft, wgt, dp, i+1, cap-wgt[i])+pft[i])
	}
	dp[i][cap] = value
	return value
}

// Dynamic dyanmic programing loop solution
func Dynamic(profits []int, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := len(profits) - 2; i >= 0; i-- {
		for j := 0; j <= capacity; j++ {
			if j >= weights[i] {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j-weights[i]]+profits[i])
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][capacity]
}

// DynamicCompress loop soulution with One-dimensional array.
func DynamicCompress(profits []int, weights []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i := 0; i < len(profits)-2; i++ {
		for j := capacity; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+profits[i])
		}
	}
	return dp[capacity]
}

// DynamicHash dyanmic programing solution with hash dictionary.
func DynamicHash(profits []int, weights []int, capacity int) int {
	dp := make([]map[int]int, len(profits))
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	return dynamicHash(profits, weights, dp, 0, capacity)
}

func dynamicHash(pft, wgt []int, dp []map[int]int, i, cap int) int {
	if val, ok := dp[i][cap]; ok {
		return val
	}
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = dynamicHash(pft, wgt, dp, i+1, cap)
	} else {
		value = max(dynamicHash(pft, wgt, dp, i+1, cap),
			dynamicHash(pft, wgt, dp, i+1, cap-wgt[i])+pft[i])
	}
	dp[i][cap] = value
	return value
}
