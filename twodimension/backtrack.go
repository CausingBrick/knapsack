package twodimension

// Backtrack solution for the two-dimensional cost knapsack problem
func Backtrack(items []*Item, kp *KnapSack) int {
	return backtrack(items, kp.Capacity, kp.Volume, 0)
}

func backtrack(items []*Item, cap, vol, i int) int {
	value := 0
	switch {
	case i == len(items):
		break
	case items[i].Bulk > vol || items[i].Weight > cap:
		value = backtrack(items, cap, vol, i+1)
	default:
		value = max(
			backtrack(items, cap, vol, i+1),
			backtrack(
				items,
				cap-items[i].Weight,
				vol-items[i].Bulk,
				i+1,
			)+items[i].Value,
		)
	}
	return value
}
