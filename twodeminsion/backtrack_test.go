package twodeminsion

import (
	"log"
	"testing"
)

type testTable struct {
	kp  *KnapSack
	its Items
}

// Add test data from here.
// Key: Name for dataset, Value: path + set name
var setPath = map[string]string{
	"set05": "./testdata/",
	"set20": "./testdata/",
	// "set40":  "./testdata/",
	// "set100": "./testdata",
	// "set500": "./testdata",
}

var testTables []*testTable

//Load data from the datasetNames befrore running the test.
func init() {
	for n, p := range setPath {
		kp, its := GetSet2D(p, n)
		testTables = append(testTables, &testTable{kp, its})
	}
}

func TestBacktrack(t *testing.T) {
	for _, v := range testTables {
		log.Println(Backtrack(v.its, v.kp))
	}
}
