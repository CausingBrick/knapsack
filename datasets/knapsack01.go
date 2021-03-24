package datasets

type KnapSack01 struct {
	Capacity   int
	Profits    []int
	Weights    []int
	selected   []int
	BestProfit int
}

// KnapSack01New returns a data set with setName form setPath
func KnapSack01New(setPath, setName string) *KnapSack01 {
	if setPath == "" || setName == "" {
		panic("Worng path for dataset")
	}

	var (
		dataset = &KnapSack01{}
		err     error
	)

	capacity, err := readToLine(setPath, setName, cSuffix)
	if err != nil {
		panic(err)
	}
	dataset.Capacity = capacity[0]

	dataset.selected, err = readToLine(setPath, setName, sSuffix)
	if err != nil {
		panic(err)
	}

	dataset.Profits, err = readToLine(setPath, setName, pSuffix)
	if err != nil {
		panic(err)
	}

	dataset.Weights, err = readToLine(setPath, setName, wSuffix)
	if err != nil {
		panic(err)
	}
	//Caculate the best profit
	for i, v := range dataset.selected {
		if v != 0 {
			dataset.BestProfit += dataset.Profits[i]
		}
	}
	if !dataset.checkBestProfit() {
		panic("Wrong selection in set :" + setName)
	}
	return dataset
}

//checkBestProfit check if the data set answer is right
func (d *KnapSack01) checkBestProfit() bool {
	weights := 0
	for i, v := range d.selected {
		if v != 0 {
			weights += d.Weights[i]
		}
	}
	return weights == d.Capacity
}
