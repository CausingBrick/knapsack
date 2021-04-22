package twodeminsion

import (
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

type dataset struct {
	name string
	path string
	kp   *KnapSack
	its  Items
	slt  int
}

func datasetNew(name, path string) *dataset {
	ds := dataset{name: name, path: path}
	ds.kp, ds.its, ds.slt = getSet2D(path, name)
	return &ds
}

var extension = ".txt"

var (
	// item profit
	prftSfx = "_p" + extension
	// item weight
	wghtSfx = "_w" + extension
	// item bulk
	bulkSfx = "_b" + extension

	// kp capacity
	capSfx = "_c" + extension
	// kp volume
	volSfx = "_v" + extension

	// soulution
	sltSfx = "_s" + extension
)

// readToLine returns all data into a singe line.
func readToLine(setPath, setName, suffix string) ([]int, error) {
	filePath := path.Join(setPath, setName+suffix)
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

// getSet2D gets data forom setpath with setName
func getSet2D(setPath, setName string) (*KnapSack, Items, int) {
	if setPath == "" || setName == "" {
		panic("Worng path for dataset")
	}

	cap, err := readToLine(setPath, setName, capSfx)
	if err != nil {
		panic(err)
	}

	vol, err := readToLine(setPath, setName, volSfx)
	if err != nil {
		panic(err)
	}

	kp := KnapSackNew(cap[0], vol[0])

	val, err := readToLine(setPath, setName, prftSfx)
	if err != nil {
		panic(err)
	}

	wght, err := readToLine(setPath, setName, wghtSfx)
	if err != nil {
		panic(err)
	}

	bulk, err := readToLine(setPath, setName, bulkSfx)
	if err != nil {
		panic(err)
	}

	its := make([]*Item, len(val))
	for i := range its {
		its[i] = ItemNew(wght[i], val[i], bulk[i])
	}

	slt, err := readToLine(setPath, setName, sltSfx)
	if err != nil {
		panic(err)
	}

	return kp, its, slt[0]
}
