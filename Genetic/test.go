package Genetic

import (
	"fmt"
	"math/rand"
)

func Zad1() {

	populationCounter := 10
	chromosomeLen := 10
	generator := func(len int) Chromosome {
		chromosome := make([]float64, len)
		for each := range chromosome {
			chromosome[each] = float64(rand.Int() % 2)
		}
		return chromosome
	}

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
	selectionFunc := func(population []Chromosome, eliteCount int, fitnessFunc func(Chromosome) float64) []Chromosome {
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
	eliteCount := 2

	// free choice with offspring of 2
	// one point crossing
	crossingFunc := func(mother Chromosome, father Chromosome) []Chromosome {
		point := rand.Intn(len(mother))
		moreLovedChild := make(Chromosome, 0, len(mother))
		moreLovedChild = append(moreLovedChild, mother[:point]...)
		moreLovedChild = append(moreLovedChild, father[point:]...)

		lessLovedChild := make(Chromosome, 0, len(mother))
		lessLovedChild = append(lessLovedChild, father[:point]...)
		lessLovedChild = append(lessLovedChild, mother[point:]...)
		return []Chromosome{moreLovedChild, lessLovedChild}
	}

	// replacement
	mutationFunc := func(chromosome Chromosome) Chromosome {
		point := rand.Intn(len(chromosome))
		if chromosome[point] == 1.0 {
			chromosome[point] = 0.0
		} else {
			chromosome[point] = 1.0
		}
		return chromosome
	}
	mutationProb := 0.6
	iterationsCap := 666

	population := GeneratePopulation(populationCounter, chromosomeLen, generator)

	// crossingProb == eliteFactor
	winnerChromosome := NaturalSelection(population,
		fitnessFunc,
		eliteCount,
		selectionFunc,
		crossingFunc,
		mutationProb,
		mutationFunc,
		iterationsCap,
		)

	fmt.Println(winnerChromosome)
}
