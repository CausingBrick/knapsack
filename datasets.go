package knapsack

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var p01 = "./datasets/p01/"

func getData(path string) ([]int, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	strArr := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var data []int
	for _, str := range strArr {
		num, err := strconv.Atoi(strings.Replace(str, " ", "", -1))
		if err != nil {
			return data, err
		}
		data = append(data, num)
	}
	return data, nil
}

func getBestprofit() int {
	profits, profit := getProfits(), 0
	selections, err := getData(p01 + "p01_s.txt")
	if err != nil {
		log.Println(err)
	}
	for i, selection := range selections {
		if selection == 1 {
			profit += profits[i]
		}
	}
	return profit
}

func getCapacity() int {
	bytes, err := ioutil.ReadFile(p01 + "p01_c.txt")

	if err != nil {
		log.Println(err)
	}

	cap, err := strconv.Atoi(strings.Replace(string(bytes), "\n", "", -1))
	if err != nil {
		log.Panic(err)
	}
	return cap
}

func getProfits() []int {
	bytes, err := ioutil.ReadFile(p01 + "p01_p.txt")

	if err != nil {
		log.Println(err)
	}
	profitsStrs := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var profits []int
	for _, profitStr := range profitsStrs {
		profitStr := profitStr
		cur, err := strconv.Atoi(profitStr)
		if err != nil {
			log.Println(err)
		}
		profits = append(profits, cur)
	}
	return profits
}

func getWeights() []int {
	bytes, err := ioutil.ReadFile(p01 + "p01_w.txt")

	if err != nil {
		log.Println(err)
	}
	weightsStrs := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var weights []int
	for _, weightsStr := range weightsStrs {
		profitStr := weightsStr
		cur, err := strconv.Atoi(profitStr)
		if err != nil {
			log.Println(err)
		}
		weights = append(weights, cur)
	}
	return weights
}
