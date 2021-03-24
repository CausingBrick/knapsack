package knapsack01

import (
	"sort"
)

type item struct {
	weight  int
	profits int
}

type sortByWeight []item

func (a sortByWeight) Len() int           { return len(a) }
func (a sortByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }

// Knapsack01Weight strategie with picking the lightest item in each selection
func Knapsack01Weight(profits []int, weights []int, capacity int) int {
	items := make(sortByWeight, len(weights))
	for i := range items {
		items[i].profits = profits[i]
		items[i].weight = weights[i]
	}
	sort.Sort(items)

	value := 0
	for _, item := range items {
		capacity -= item.weight
		if capacity < 0 {
			break
		}
		value += item.profits
	}
	return value
}

type sortByProfit []item

func (a sortByProfit) Len() int           { return len(a) }
func (a sortByProfit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByProfit) Less(i, j int) bool { return a[i].profits > a[j].profits }

// Knapsack01Profit Pick the best profit item in each selection.
func Knapsack01Profit(profits []int, weights []int, capacity int) int {
	items := make(sortByProfit, len(profits))
	for i := range items {
		items[i].profits = profits[i]
		items[i].weight = weights[i]
	}
	sort.Sort(items)
	value := 0
	for _, item := range items {
		capacity -= item.weight
		if capacity < 0 {
			break
		}
		value += item.profits
	}
	return value
}

type sortByRatio []item

func (a sortByRatio) Len() int      { return len(a) }
func (a sortByRatio) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByRatio) Less(i, j int) bool {
	ratioi := float32(a[i].profits) / float32(a[i].weight)
	ratioj := float32(a[j].profits) / float32(a[j].weight)
	return ratioi > ratioj
}

// Knapsack01Ratio Pick the item has highest ratio of profit and weight in each selection
func Knapsack01Ratio(profits []int, weights []int, capacity int) int {
	items := make(sortByRatio, len(profits))
	for i := range items {
		items[i].profits = profits[i]
		items[i].weight = weights[i]
	}
	sort.Sort(items)

	value := 0
	for _, item := range items {
		capacity -= item.weight
		if capacity < 0 {
			break
		}
		value += item.profits
	}
	return value
}

// Knapsack01Recursive recursve solution
func Knapsack01Recursive(profits []int, weights []int, capacity int) int {
	return knapsack01Recursive(profits, weights, 0, capacity)
}

func knapsack01Recursive(pft, wgt []int, i, cap int) int {
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = knapsack01Recursive(pft, wgt, i+1, cap)
	} else {
		value = max(knapsack01Recursive(pft, wgt, i+1, cap),
			knapsack01Recursive(pft, wgt, i+1, cap-wgt[i])+pft[i])
	}
	return value
}

// Knapsack01Dynamic dyanmic programing recursive solution
func Knapsack01Dynamic(profits []int, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	return knapsack01Dynamic(profits, weights, dp, 0, capacity)
}

func knapsack01Dynamic(pft, wgt []int, dp [][]int, i, cap int) int {
	if dp[i][cap] > 0 {
		return dp[i][cap]
	}
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = knapsack01Dynamic(pft, wgt, dp, i+1, cap)
	} else {
		value = max(knapsack01Dynamic(pft, wgt, dp, i+1, cap),
			knapsack01Dynamic(pft, wgt, dp, i+1, cap-wgt[i])+pft[i])
	}
	dp[i][cap] = value
	return value
}

// Knapsack01DynamicLoop dyanmic programing loop solution
func Knapsack01DynamicLoop(profits []int, weights []int, capacity int) int {
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

// Knapsack01DynamicLoopOpt loop soulution with One-dimensional array.
func Knapsack01DynamicLoopOpt(profits []int, weights []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i := 0; i < len(profits)-2; i++ {
		for j := capacity; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+profits[i])
		}
	}
	return dp[capacity]
}

// Knapsack01DynamicHash dyanmic programing solution with hash dictionary.
func Knapsack01DynamicHash(profits []int, weights []int, capacity int) int {
	dp := make([]map[int]int, len(profits))
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	return knapsack01DynamicHash(profits, weights, dp, 0, capacity)
}

func knapsack01DynamicHash(pft, wgt []int, dp []map[int]int, i, cap int) int {
	if val, ok := dp[i][cap]; ok {
		return val
	}
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = knapsack01DynamicHash(pft, wgt, dp, i+1, cap)
	} else {
		value = max(knapsack01DynamicHash(pft, wgt, dp, i+1, cap),
			knapsack01DynamicHash(pft, wgt, dp, i+1, cap-wgt[i])+pft[i])
	}
	dp[i][cap] = value
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
