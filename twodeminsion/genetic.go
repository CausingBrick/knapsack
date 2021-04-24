package twodeminsion

import (
	"math/rand"
	"sync"
)

// GenConf defines the config for SGA
type GenConf struct {
	// size of population
	PopulationSize int
	// Competition Scale for tournament Selection
	TournamentSize int
	// Number of evolutions
	Generations int
	//  Probability of Chromosome Variation
	mutateRatio float64
}

type individual []bool

func (ind individual) fitness(its Items, kp *KnapSack) int {
	var sumVal, sumWght, sumBulk int
	for i, selected := range ind {
		if selected {
			sumVal += its[i].Value
			sumWght += its[i].Weight
			sumBulk += its[i].Bulk
		}
	}
	if kp.Capacity < sumWght || kp.Volume < sumBulk {
		sumVal = 0
	}
	return sumVal
}

func (ind individual) mutation(mutationRatio float32) {
	if mutationRatio > rand.Float32() {
		rndIndex := rand.Intn(len(ind))
		(ind)[rndIndex] = !(ind)[rndIndex]
		rndIndex = rand.Intn(len(ind))
		(ind)[rndIndex] = !(ind)[rndIndex]
	}
}

type population []individual

func (pop *population) fitnesses(items Items, kp *KnapSack) []int {
	fitness := make([]int, len(*pop))
	var wg sync.WaitGroup
	wg.Add(len(*pop))
	for i := range *pop {
		go func(i int) {
			fitness[i] = (*pop)[i].fitness(items, kp)
			wg.Done()
		}(i)

	}
	wg.Wait()
	return fitness
}

func (pop *population) removeWeakest(items Items, kp *KnapSack) {
	fitness := pop.fitnesses(items, kp)
	minIndex := 0
	for i, f := range fitness {
		if f < fitness[minIndex] {
			minIndex = i
		}
	}
	*pop = append((*pop)[:minIndex], (*pop)[minIndex+1:]...)
}

func (pop *population) Add(ind individual) {
	*pop = append(*pop, ind)
}

// Genetic simple genetic algorithm.
func Genetic(items Items, kp *KnapSack, conf *GenConf) int {
	maxv := 0
	pop := populationInit(len(items), conf.PopulationSize)
	for i := 0; i < conf.Generations; i++ {
		indA := tournamentSelection(pop, conf.TournamentSize, items, kp)
		indB := tournamentSelection(pop, conf.TournamentSize, items, kp)

		childA, childB := crossover(indA, indB)
		childA.mutation(float32(conf.mutateRatio))
		childB.mutation(float32(conf.mutateRatio))

		pop.removeWeakest(items, kp)
		pop.Add(childA)
		pop.removeWeakest(items, kp)
		pop.Add(childB)
		if cur := max(pop.fitnesses(items, kp)...); maxv < cur {
			maxv = cur
		}
	}

	return maxv
}

func populationInit(itemsNum, populationSize int) population {
	pop := make([]individual, populationSize)
	var wg sync.WaitGroup
	wg.Add(populationSize)
	for i := range pop {
		go func(i int) {
			pop[i] = make(individual, itemsNum)
			for j := range pop[i] {
				pop[i][j] = rand.Int31n(2) == 1
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return pop
}

func tournamentSelection(pop population, tournamentSize int, items Items, kp *KnapSack) individual {
	var wg sync.WaitGroup
	wg.Add(tournamentSize)
	fitness := make([]int, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		go func(i int) {
			rndIndex := rand.Intn(len(pop))
			fitness[i] = pop[rndIndex].fitness(items, kp)
			wg.Done()
		}(i)
	}
	wg.Wait()

	maxIndex := 0
	for i, f := range fitness {
		if f > fitness[maxIndex] {
			maxIndex = i
		}
	}
	return pop[maxIndex]
}

func crossover(parentA, parentB individual) (individual, individual) {
	axis := rand.Intn(len(parentA)-1) + 1
	childA := append(parentA[:axis], parentB[axis:]...)
	childB := append(parentB[:axis], parentA[axis:]...)
	return childA, childB
}
