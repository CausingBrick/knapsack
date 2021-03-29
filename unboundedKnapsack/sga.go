package unboundedKnapsack

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/CausingBrick/knapsack"
)

// SGAConf defines the config for SGA
type SGAConf struct {
	// size of population
	PopulationSize int
	// Competition Scale for tournament Selection
	CompetitionScale int
	// Number of evolutions
	Generations int
	//  Probability of Chromosome Variation
	mutateRatio float64
}

type chromosome []int

type individual struct {
	chromo        chromosome
	fitnessProfit int
	fitnessWeight int
}

func (ind *individual) caculateFitness(items []*knapsack.Item) {
	ind.fitnessProfit, ind.fitnessWeight = 0, 0
	for i, quantity := range ind.chromo {
		ind.fitnessProfit += quantity * items[i].Value
		ind.fitnessWeight += quantity * items[i].Weight
	}
}

type parents struct {
	mother, father *individual
}

type population []*individual

// SGA simple genetic algorithm for ukp
func SGA(items []*knapsack.Item, capacity int, conf *SGAConf) []int {
	//Get the real encoding
	chrom := chromosomeEncode(items, capacity)
	// Get the first generation population
	pop := populationInit(items, capacity, chrom, conf.PopulationSize)
	profits := make([]int, conf.Generations)
	// Enter envolution
	for i := 0; i < conf.Generations; i++ {
		pop = realValuedCrossover(items, capacity, pop, conf.CompetitionScale, conf.PopulationSize, len(chrom))
		realValuedMutation(items, capacity, pop, conf.mutateRatio)
		// Select the best indivadual
		for _, ind := range pop {
			if profits[i] < ind.fitnessProfit {
				profits[i] = ind.fitnessProfit
			}
		}
	}
	return profits
}

// chromomeEncode encodes the number of each item that the backpack can
// hold into a binary string of dynamic length.
func chromosomeEncode(items []*knapsack.Item, capacity int) chromosome {
	chromosome := make([]int, len(items))
	for i, item := range items {
		chromosome[i] = capacity/item.Weight + 1
	}
	return chromosome
}

// populationInit returns an intialized population
func populationInit(items []*knapsack.Item, capacity int, chromosome []int, populationSize int) population {
	population := *new(population)
	for i := 0; i < populationSize; {
		ind := &individual{
			make([]int, len(chromosome)),
			0, 0,
		}
		for i, quantity := range chromosome {
			randChromosome := rand.Intn(quantity)
			// fmt.Println(quantity, randChromosome)
			ind.chromo[i] = randChromosome
			ind.fitnessProfit += randChromosome * items[i].Value
			ind.fitnessWeight += randChromosome * items[i].Weight
		}
		// Check whether the individual's weight is qualified
		if ind.fitnessWeight > capacity {
			continue
		} else {
			i++
			population = append(population, ind)
		}
	}
	return population
}

func (a population) Len() int           { return len(a) }
func (a population) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a population) Less(i, j int) bool { return a[i].fitnessProfit > a[j].fitnessProfit }

//rankingSelection returns the parents from the src int the range between 0 and seletionRatio
// TODO: Use concurrency optimization
func rankingSelection(src population, selectionRatio float32) parents {
	populationSize := len(src)
	selectionRange := populationSize * int(selectionRatio*1024) >> 10
	sort.Sort(src)
	selection := func() *individual {
		return src[rand.Intn(selectionRange)]
	}
	return parents{selection(), selection()}
}

//tournamentSelection returns the parents from the src
// TODO: Use concurrency optimization
func tournamentSelection(src population, competitionScale int) parents {
	populationSize := len(src)
	//To save iteration times
	if competitionScale > populationSize {
		competitionScale = populationSize
	}
	selection := func() *individual {
		var max *individual
		for i := 0; i < competitionScale; i++ {
			n := rand.Intn(populationSize)
			rander := src[n]
			if max != nil {
				if max.fitnessProfit < rander.fitnessProfit {
					max = rander
				}
			} else {
				max = rander
			}
		}
		return max
	}
	return parents{selection(), selection()}
}

// realValuedCrossover returns the childern population after crossovered
func realValuedCrossover(items []*knapsack.Item, capacity int, src population, competitionScale, populationSize, chromLen int) population {
	envolutionPop := *new([]*individual)
	// Crossover
	for popSize := 0; popSize < populationSize; {
		// Selection individual to be parents
		parent := tournamentSelection(src, competitionScale)
		// parent := rankingSelection(src, 0.6)
		//Crossover points range from 2 to n-1
		division := rand.Intn(chromLen-2) + 2
		boy := &individual{}
		boy.chromo = append(boy.chromo, parent.mother.chromo[:division]...)
		boy.chromo = append(boy.chromo, parent.father.chromo[division:]...)
		// Check whether the individual's weight is qualified
		boy.caculateFitness(items)
		if boy.fitnessWeight <= capacity {
			envolutionPop = append(envolutionPop, boy)
			popSize++
			if popSize == populationSize {
				break
			}
		}
		girl := &individual{}
		girl.chromo = append(girl.chromo, parent.father.chromo[:division]...)
		girl.chromo = append(girl.chromo, parent.mother.chromo[division:]...)
		// Check whether the individual's weight is qualified
		girl.caculateFitness(items)
		if girl.fitnessWeight <= capacity {
			popSize++
			envolutionPop = append(envolutionPop, girl)
			if popSize == populationSize {
				break
			}
		}
	}
	return envolutionPop
}

// realValuedMutation returns the childern population after mutation
func realValuedMutation(items []*knapsack.Item, capacity int,
	pop population, mutationRatio float64) {
	if mutationRatio > rand.Float64() {
		mutatedInd := rand.Intn(len(pop))
		chromLen := len(pop[mutatedInd].chromo)
		for {
			// Generate the index of the mutation location
			mutateX, mutateY := rand.Intn(chromLen), rand.Intn(chromLen)
			for mutateX == mutateY {
				mutateX, mutateY = rand.Intn(chromLen), rand.Intn(chromLen)
			}
			mutatedChrom := make(chromosome, chromLen)
			copy(mutatedChrom, pop[mutatedInd].chromo)
			//Mutated the chromosome
			mutatedChrom[mutateX], mutatedChrom[mutateY] =
				mutatedChrom[mutateY], mutatedChrom[mutateX]

			//Check whether the chromosome is a good variant
			quantity := 0
			for i, item := range items {
				quantity += item.Weight * mutatedChrom[i]
			}
			// Chromosome qualified
			if quantity <= capacity {
				pop[mutatedInd].chromo = mutatedChrom
				break
			}
		}
	}
}

func printPop(p population) {
	for _, v := range p {
		fmt.Print(v)
	}
	fmt.Println()
}

// spCrossover single-point crossover
// func spCrossover(parents []parents, populationSize int) population {
// 	population := make([]*individual, populationSize)

// 	for i := 0; i < populationSize; {
// 		ind := new(individual)
// 		for j, parent := range parents {
// 			chromosomeX := parent.mother.chromosome
// 			chromosomeY := parent.father.chromosome
// 			var chromosomeA, chrchromosomeB chromosome
// 			for k := range chromosomeX {
// 				division := knapsack.Max(bitLen(chromosomeX[k]), bitLen(chromosomeY[k])) >> 1
// 				chromosomeA[k] = (1 << division) - 1
// 				chromosomeA[k] &=
// 			}
// 		}
// 		// Check whether the individual's weight is qualified
// 		if ind.fitnessWeight > capacity {
// 			continue
// 		} else {
// 			i++
// 			population = append(population, ind)
// 		}
// 	}
// 	var boy, girl *individual
// 	for i := 0; i < populationSize; i++ {
// 		for j := 0; j < parents[i].father.chromosome; j++ {

// 		}
// 	}
// }

