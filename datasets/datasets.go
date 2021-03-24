package datasets

import (
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

var (
	extension = ".txt"
	cSuffix   = "_c" + extension
	pSuffix   = "_p" + extension
	wSuffix   = "_w" + extension
	sSuffix   = "_s" + extension
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
