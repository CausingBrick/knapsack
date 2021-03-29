package knapsack

type Item struct {
	Weight int
	Value  int
}

func ItemNew(w, v int) *Item {
	return &Item{w, v}
}

func ItemsNew(ws, vs []int) []*Item {
	if len(ws) != len(vs) {
		return nil
	}
	items := make([]*Item, len(ws))
	for i := range items {
		items[i] = ItemNew(ws[i], vs[i])
	}
	return items
}

type KnapSack struct {
	Items    []*Item
	Capacity int
}

func New() *KnapSack {
	return &KnapSack{make([]*Item, 0), 0}
}
