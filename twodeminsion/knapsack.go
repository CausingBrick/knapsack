package twodeminsion

type KnapSack struct {
	Capacity int
	Volume   int
}

func KnapSackNew(capacity, volume int) *KnapSack {
	return &KnapSack{
		Capacity: capacity,
		Volume:   volume,
	}
}

type Item struct {
	Weight int
	Value  int
	Bulk   int
}

func ItemNew(weight, value, bulk int) *Item {
	return &Item{
		Weight: weight,
		Value:  value,
		Bulk:   bulk,
	}
}

type Items []*Item
