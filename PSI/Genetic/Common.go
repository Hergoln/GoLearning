package Genetic

import (
	"math/rand"
)

type Chromosome = []float64
type selector func([]Chromosome, int, []float64) []Chromosome
type Type = interface{}

func GeneratePopulation(n, chromosomeLen int, generator func(int) Chromosome) []Chromosome {
	population := make([]Chromosome, n)
	for each := range population {
		population[each] = generator(chromosomeLen)
	}

	return population
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

func sortChromosomesByFitDesc(population []Chromosome, fitValues []float64) []Chromosome {
	var fit, max float64
	ind := 0
	populationCopy := make([]Chromosome, 0, len(population))
	for each := range population {
		populationCopy = append(populationCopy, population[each])
	}
	var elites []Chromosome
	for vacant := 0; len(populationCopy) > 0; vacant-- {
		max = fitValues[0]
		fit = max
		ind = 0

		for each := range populationCopy {
			fit = fitValues[each]
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

func deepCopy(toCopy []Chromosome) []Chromosome {
	toRet := make([]Chromosome, len(toCopy))

	for each := range toRet {
		toRet[each] = toCopy[each]
	}

	return toRet
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

func simpleGenerator(chromLen int) Chromosome {
	chromosome := make([]float64, chromLen)
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