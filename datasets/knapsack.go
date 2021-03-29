package datasets

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/CausingBrick/knapsack"
)

// ReadDataset returns a data set with setName form setPath
func ReadDataset(setPath string) *knapsack.KnapSack {

	kp := knapsack.New()
	// read capacity
	capacity, _ := ReadToLine(setPath + cSuffix)
	kp.Capacity = capacity[0]
	// read value and weight
	values, _ := ReadToLine(setPath + pSuffix)
	weights, _ := ReadToLine(setPath + wSuffix)
	kp.Items = knapsack.ItemsNew(values, weights)
	return kp
}

// readToLine returns all data into a singe line.
func ReadToLine(filePath string) ([]int, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	//Split into rows after removes extra spaces
	rows := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var data []int
	for _, str := range rows {
		num, err := strconv.Atoi(strings.Replace(str, " ", "", -1))
		if err != nil {
			return data, err
		}
		data = append(data, num)
	}
	return data, nil
}
