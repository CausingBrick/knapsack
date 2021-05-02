package twodimensionpacking

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Design for reading dataset CGCUTBIN
type dataset struct {
	name string
	path string
	its  []*Item
}

func datasetNew(name, filePath string) *dataset {
	lines, err := readToLines(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	// boxH, boxW := lines[1], lines[2]
	items := []*Item{}
	for _, line := range lines {
		items = append(items, ItemNew(line[0], line[1]))
	}
	ds := &dataset{
		name: name,
		path: filePath,
		// box:  BoxNew(boxH, boxW),
		its: items,
	}
	return ds
}

func readToLines(filePath string) ([][]int, error) {
	bts, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(bts)), "\n")
	nums := make([][]int, len(lines))
	for i, line := range lines {
		chars := strings.Split(strings.TrimSpace(line), " ")
		for _, char := range chars {
			num, _ := strconv.Atoi(char)
			nums[i] = append(nums[i], num)
		}
	}
	return nums, nil
}
