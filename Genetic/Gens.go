package Genetic

import (
	"math/rand"
	"time"
)

type Chromosome = []float64
type selector func([]Chromosome, int, func (Chromosome) float64) []Chromosome
type Type = interface{}
// only for float64 temporarily
func GeneratePopulation(n, chromosomeLen int, generator func(int) Chromosome) []Chromosome {
	population := make([]Chromosome, n)
	for each := range population {
		population[each] = generator(chromosomeLen)
	}

	return population
}

// iteration steps:
// Selection -> Crossing -> Mutation
func NaturalSelection(population []Chromosome,
	fitnessFunc func (Chromosome) float64,
	eliteCount int,
	selectionFunc selector,
	crossingFunc func(Chromosome, Chromosome) []Chromosome,
	mutationProb float64,
	mutationFunc func(Chromosome) Chromosome,
	iterationsCap int,
	) Chromosome {
	rand.Seed(time.Now().UnixNano())

	fitValues := make([]float64, len(population))
	var offspring, elites []Chromosome


	// main algorithm loop
	for iter := 0; iter < iterationsCap; iter++ {
		for each := range population {
			fitValues[each] = fitnessFunc(population[each])
			if fitValues[each] == 1.0 {
				return population[each]
			}
		}

		// new population
		elites = selectionFunc(population, eliteCount, fitnessFunc)
		for each := range elites {
			probValue := rand.Float64()
			if probValue < mutationProb {
				elites[each] = mutationFunc(elites[each])
			}
		}

		offspring = crossingFunc(elites[0], elites[1])

		population = sortChromosomesByFitDesc(population, fitnessFunc)
		for each := range offspring {
			population[len(population) - 1 - each] = offspring[each]
		}
	}

	// end
	return maxFromPopulation(population, fitnessFunc)
}

func maxFromPopulation(population []Chromosome, fitnessFunc func (Chromosome) float64) Chromosome {
	fitValues := make([]float64, len(population))
	max := fitnessFunc(population[0])
	ind := 0
	for each := range population {
		fitValues[each] = fitnessFunc(population[each])
		if max < fitValues[each] {
			max = fitValues[each]
			ind = each
		}
	}

	return population[ind]
}

func sortChromosomesByFitDesc(population []Chromosome, fitnessFunc func (Chromosome) float64) []Chromosome {
	var fit, max float64
	ind := 0
	populationCopy := make([]Chromosome, 0, len(population))
	for each := range population {
		populationCopy = append(populationCopy, population[each])
	}
	var elites []Chromosome
	for vacant := 0; len(populationCopy) > 0; vacant-- {
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
