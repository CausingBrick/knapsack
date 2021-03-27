package knapsack01

import "sort"

type weightSorter []Item

func (a weightSorter) Len() int           { return len(a) }
func (a weightSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a weightSorter) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

// GreedyWeight strategie with picking the lightest item in each selection
func GreedyWeight(profits []int, weights []int, capacity int) int {
	items := make(weightSorter, len(weights))
	for i := range items {
		items[i].Value = profits[i]
		items[i].Weight = weights[i]
	}
	sort.Sort(items)

	value := 0
	for _, item := range items {
		capacity -= item.Weight
		if capacity < 0 {
			break
		}
		value += item.Value
	}
	return value
}

type valueSorter []Item

func (a valueSorter) Len() int           { return len(a) }
func (a valueSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a valueSorter) Less(i, j int) bool { return a[i].Value > a[j].Value }

// Knapsack01Profit Pick the best profit item in each selection.
func Knapsack01Profit(profits []int, weights []int, capacity int) int {
	items := make(valueSorter, len(profits))
	for i := range items {
		items[i].Value = profits[i]
		items[i].Weight = weights[i]
	}
	sort.Sort(items)
	value := 0
	for _, item := range items {
		capacity -= item.Weight
		if capacity < 0 {
			break
		}
		value += item.Value
	}
	return value
}

type ratioSorter []Item

func (a ratioSorter) Len() int      { return len(a) }
func (a ratioSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ratioSorter) Less(i, j int) bool {
	ratioi := float32(a[i].Value) / float32(a[i].Weight)
	ratioj := float32(a[j].Value) / float32(a[j].Weight)
	return ratioi > ratioj
}

// Knapsack01Ratio Pick the item has highest ratio of profit and weight in each selection
func Knapsack01Ratio(profits []int, weights []int, capacity int) int {
	items := make(ratioSorter, len(profits))
	for i := range items {
		items[i].Value = profits[i]
		items[i].Weight = weights[i]
	}
	sort.Sort(items)

	value := 0
	for _, item := range items {
		capacity -= item.Weight
		if capacity < 0 {
			break
		}
		value += item.Value
	}
	return value
}
