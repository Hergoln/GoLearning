package Genetic

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Zad1() {
	populationCounter := 10
	chromosomeLen := 10

	// this function will return %-wise fitness level
	fitnessFunc := func(chromosome Chromosome) float64 {
		if len(chromosome) == 0 {
			return 0.0
		}
		sum := 0.0
		for each := range chromosome {
			sum += chromosome[each]
		}
		return sum / float64(len(chromosome))
	}

	// rank selection
	selectionFunc := func(population []Chromosome,
		eliteCount int,
		fitnessFunc func(Chromosome) float64,
		) []Chromosome {
		var fit, max float64
		ind := 0
		populationCopy := make([]Chromosome, 0, len(population))
		for each := range population {
			populationCopy = append(populationCopy, population[each])
		}
		var elites []Chromosome
		for vacant := eliteCount; vacant > 0 && len(populationCopy) > 0; vacant-- {
			max = fitnessFunc(populationCopy[0])
			fit = max
			ind = 0

			for each := range populationCopy {
				fit = fitnessFunc(populationCopy[each])
				if max < fit {
					max = fit
					ind = each
				}
			}

			elites = append(elites, populationCopy[ind])
			populationCopy = append(populationCopy[0:ind], populationCopy[ind+1:]...)
		}

		return elites
	}
	eliteCounter := 2

	// free choice with offspring of 2
	crossingFunc := onePointCrossing

	mutationFunc := replacementMutation
	mutationProb := 0.6
	iterationsCap := 64

	population := GeneratePopulation(populationCounter, chromosomeLen, simpleGenerator)

	// crossingProb == eliteFactor
	winnerChromosome := SimpleGens(population,
		fitnessFunc,
		eliteCounter,
		selectionFunc,
		crossingFunc,
		mutationProb,
		mutationFunc,
		iterationsCap,
		)

	fmt.Println(winnerChromosome)
}

func Zad2() {

	populationCounter := 10
	chromosomeLen := 8

	fitnessFunc := func(chromosome Chromosome) float64 {
		A, B := getABFromChromosome(chromosome)
		res := math.Abs(2.0 * A * A + B - 33.0) + 1
		return 1 / res
	}

	// this roulette might choose same chromosome twice
	selectionFunc := func(population []Chromosome,
		eliteCounter int,
		fitnessFunc func(Chromosome) float64,
		) []Chromosome  {
		rand.Seed(time.Now().UnixNano())

		elites := make([]Chromosome, 0)

		fitValues := make([]float64, len(population))
		valuesMap := make(map[float64]Chromosome)
		prevMapKey, sum:= 0.0, 0.0

		for each := range fitValues {
			fitValues[each] = math.Abs(fitnessFunc(population[each]))
			sum += fitValues[each]
		}

		for each := range fitValues {
			prevMapKey += fitValues[each] / sum
			valuesMap[prevMapKey] = population[each]
		}

		for i := 0; i < eliteCounter; i++ {
			rul := roulette(valuesMap)
			elites = append(elites, rul)
		}

		return elites
	}
	eliteCounter := populationCounter

	// free choice with offspring of 2
	crossingFunc := onePointCrossing

	mutationFunc := replacementMutation
	// those parameters give good success ratio
	mutationProb := 0.5
	iterationsCap := 512

	population := GeneratePopulation(populationCounter, chromosomeLen, simpleGenerator)

	winnerChromosome := FindEquationSolution(population,
		fitnessFunc,
		eliteCounter,
		selectionFunc,
		crossingFunc,
		mutationProb,
		mutationFunc,
		iterationsCap,
	)

	fmt.Println(getABFromChromosome(winnerChromosome))
}

func simpleGenerator(len int) Chromosome {
	chromosome := make([]float64, len)
	for each := range chromosome {
		chromosome[each] = float64(rand.Int() % 2)
	}
	return chromosome
}

func replacementMutation(chromosome Chromosome) Chromosome {
	point := rand.Intn(len(chromosome))
	if chromosome[point] == 1.0 {
		chromosome[point] = 0.0
	} else {
		chromosome[point] = 1.0
	}
	return chromosome
}

func onePointCrossing(mother Chromosome, father Chromosome) []Chromosome {
	point := rand.Intn(len(mother))
	moreLovedChild := make(Chromosome, 0, len(mother))
	moreLovedChild = append(moreLovedChild, mother[:point]...)
	moreLovedChild = append(moreLovedChild, father[point:]...)

	lessLovedChild := make(Chromosome, 0, len(mother))
	lessLovedChild = append(lessLovedChild, father[:point]...)
	lessLovedChild = append(lessLovedChild, mother[point:]...)
	return []Chromosome{moreLovedChild, lessLovedChild}
}

func getABFromChromosome(chromosome Chromosome) (float64, float64) {
	bitsA := make([]float64, len(chromosome) / 2)
	bitsB := make([]float64, len(chromosome) / 2)
	A, B := 0.0, 0.0

	for i := 0; i < len(chromosome) / 2; i++ {
		bitsA[i] = chromosome[i]
	}

	for i := 0; i < len(chromosome) / 2; i++ {
		bitsB[i] = chromosome[i + len(chromosome) / 2]
	}

	for each := range bitsA {
		A += bitsA[each] * math.Pow(float64(each), 2.)
	}

	for each := range bitsB {
		B += bitsB[each] * math.Pow(float64(each), 2.)
	}

	return A, B
}

func roulette(chromosomeMap map[float64]Chromosome) Chromosome {
	randVal := rand.Float64()
	lastKey := 0.0
	for key := range chromosomeMap {
		if randVal < key {
			return chromosomeMap[key]
		}
		lastKey = key
	}
	return chromosomeMap[lastKey]
}