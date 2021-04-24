package twodeminsion

import (
	"testing"
)

var geneticSets = map[string]*dataset{
	"set05":  datasetNew("set05", "./testdata/"),
	"set20":  datasetNew("set20", "./testdata/"),
	"set40":  datasetNew("set40", "./testdata/"),
	"set100": datasetNew("set100", "./testdata/"),
	"set250": datasetNew("set250", "./testdata/"),
	"set500": datasetNew("set500", "./testdata/"),
}
var conf = &GenConf{
	PopulationSize: 10,
	TournamentSize: 2,
	Generations:    100,
	mutateRatio:    0.618,
}

func TestGenetic05(t *testing.T) {
	got := Genetic(geneticSets["set05"].its, geneticSets["set05"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set05"].name, geneticSets["set05"].slt, got)
}
func TestGenetic20(t *testing.T) {
	got := Genetic(geneticSets["set20"].its, geneticSets["set20"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set20"].name, geneticSets["set20"].slt, got)
}

func TestGenetic40(t *testing.T) {
	got := Genetic(geneticSets["set40"].its, geneticSets["set40"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set40"].name, geneticSets["set40"].slt, got)
}

func TestGenetic100(t *testing.T) {
	got := Genetic(geneticSets["set100"].its, geneticSets["set100"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set100"].name, geneticSets["set100"].slt, got)
}
func TestGenetic250(t *testing.T) {
	got := Genetic(geneticSets["set250"].its, geneticSets["set250"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set250"].name, geneticSets["set250"].slt, got)
}
func TestGenetic500(t *testing.T) {
	got := Genetic(geneticSets["set500"].its, geneticSets["set500"].kp, conf)
	t.Logf("the %s exact value is %d, got %d", geneticSets["set500"].name, geneticSets["set500"].slt, got)
}
