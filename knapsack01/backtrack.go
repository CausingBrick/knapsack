package knapsack01

// Backtrack recursve solution
func Backtrack(profits []int, weights []int, capacity int) int {
	return backtrack(profits, weights, 0, capacity)
}

func backtrack(pft, wgt []int, i, cap int) int {
	value := 0
	if i == len(pft)-1 {
		value = 0
	} else if cap < wgt[i] {
		value = backtrack(pft, wgt, i+1, cap)
	} else {
		value = max(backtrack(pft, wgt, i+1, cap),
			backtrack(pft, wgt, i+1, cap-wgt[i])+pft[i])
	}
	return value
}
