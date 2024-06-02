package Genetic

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Pair struct {
	x float64
	y float64
}

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
	selectionFunc := eliteSelectorFunc
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
	selectionFunc := eliteSelectorFunc
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

// backpack problem
func Zad3() {
	populationCounter := 8
	chromosomeLen := 10
	iterationsCap := 2048
	eliteFactor := 0.2
	mutationProb := 0.05

	population := GeneratePopulation(populationCounter, chromosomeLen, simpleGenerator)

	winnerChromosome := backpackSolution(population, backpackProblemFitness, eliteFactor, onePointCrossing, mutationProb, eachGenMutation, iterationsCap)

	fmt.Println(getWeightsAndValues(winnerChromosome))
	for each := range winnerChromosome {
		if winnerChromosome[each] == 1.0 {
			fmt.Printf("%d ", each)
		}
	}
}

func Zad4() {
	populationCounter := 100
	chromosomeLen := 25 // id's of cities (coords)
	iterationsCap := 2048
	eliteFactor := 0.2
	mutationProb := 0.01

	population := GeneratePopulation(populationCounter, chromosomeLen, travellerGenerator)

	winnerChromosome := backpackSolution(population, travellerProblemFitness, eliteFactor, orderCrossing, mutationProb, swapNeighbourMutation, iterationsCap)

	fmt.Print(winnerChromosome)
	fmt.Printf(" %f\n", 1 / travellerProblemFitness(winnerChromosome))
}

// iteration steps:
// Selection -> Crossing -> Mutation
func SimpleGens(population []Chromosome,
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
		elites = selectionFunc(population, eliteCount, fitValues)
		for each := range elites {
			probValue := rand.Float64()
			if probValue < mutationProb {
				elites[each] = mutationFunc(elites[each])
			}
		}

		offspring = crossingFunc(elites[0], elites[1])

		population = sortChromosomesByFitDesc(population, fitValues)
		for each := range offspring {
			population[len(population) - 1 - each] = offspring[each]
		}
	}

	// end
	return maxFromPopulation(population, fitnessFunc)
}

func FindEquationSolution(population []Chromosome,
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
	//var elites []Chromosome
	// main algorithm loop
	for iter := 0; iter < iterationsCap; iter++ {
		for each := range population {
			fitValues[each] = fitnessFunc(population[each])
			if fitValues[each] == 1.0 {
				return population[each]
			}
		}

		// new population
		population = selectionFunc(population, eliteCount, fitValues)

		tempPopulation := deepCopy(population[:len(population) / 2])
		offspring := make([]Chromosome, len(tempPopulation))
		for i := 0; i < len(population) / 2; i++ {
			offspring[i] = crossingFunc(tempPopulation[i], tempPopulation[len(offspring) - 1 - i])[0]
		}

		// I believe this one only copies address(reference)
		population = deepCopy(append(population[len(population) / 2:], offspring...))
		for each := range population {
			probValue := rand.Float64()
			if probValue < mutationProb {
				population[each] = mutationFunc(population[each])
			}
		}
	}

	return maxFromPopulation(population, fitnessFunc)
}

func getABFromChromosome(chromosome Chromosome) (float64, float64) {
	bitsA := make([]float64, len(chromosome)/2)
	bitsB := make([]float64, len(chromosome)/2)
	A, B := 0.0, 0.0

	for i := 0; i < len(chromosome)/2; i++ {
		bitsA[i] = chromosome[i]
	}

	for i := 0; i < len(chromosome)/2; i++ {
		bitsB[i] = chromosome[i+len(chromosome)/2]
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
	for key := range chromosomeMap {
		if randVal < key {
			return chromosomeMap[key]
		}
	}

	counter := rand.Intn(len(chromosomeMap))
	for _, val := range chromosomeMap {
		if counter == 0 {
			return val
		}
		counter++
	}

	return nil
}

func backpackProblemFitness(chromosome Chromosome) float64 {
	values, weight := getWeightsAndValues(chromosome)
	if weight > 35 {
		return 0
	}

	return values
}

func rouletteSelectorFunc(population []Chromosome,
	eliteCounter int,
	fitValuesIn []float64,
) []Chromosome {
	rand.Seed(time.Now().UnixNano())

	elites := make([]Chromosome, 0)

	fitValues := make([]float64, len(population))
	valuesMap := make(map[float64]Chromosome)
	prevMapKey, sum := 0.0, 0.0

	for each := range fitValues {
		fitValues[each] = math.Abs(fitValuesIn[each])
		sum += fitValues[each]
	}

	for each := range fitValues {
		if sum == 0 {
			prevMapKey = 0
		} else {
			prevMapKey += fitValues[each] / sum
		}
		valuesMap[prevMapKey] = population[each]
	}

	for i := 0; i < eliteCounter; i++ {
		rul := roulette(valuesMap)
		elites = append(elites, rul)
	}

	return elites
}

func eliteSelectorFunc(population []Chromosome,
	eliteCount int,
	fitValues []float64,
) []Chromosome {
	var fit, max float64
	ind := 0
	populationCopy := make([]Chromosome, 0, len(population))
	for each := range population {
		populationCopy = append(populationCopy, population[each])
	}
	var elites []Chromosome
	for vacant := eliteCount; vacant > 0 && len(populationCopy) > 0; vacant-- {
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

func getWeightsAndValues(chromosome Chromosome) (float64, float64) {
	values := []float64{266, 442, 671, 526, 388, 245, 210, 145, 126, 322}
	weights := []float64{3, 13, 10, 9, 7, 1, 8, 8, 2, 9}

	sumWeight := 0.0
	sumValues := 0.0
	for each := range chromosome {
		if chromosome[each] == 1.0 {
			sumWeight += weights[each]
			sumValues += values[each]
		}
	}

	return sumValues, sumWeight
}

func backpackSolution(population []Chromosome,
	fitnessFunc func (Chromosome) float64,
	eliteFactor float64,
	crossingFunc func(Chromosome, Chromosome) []Chromosome,
	mutationProb float64,
	mutationFunc func(Chromosome, float64) Chromosome,
	iterationsCap int,
	) Chromosome {
	rand.Seed(time.Now().UnixNano())
	fitValues := make([]float64, len(population))
	for iter := 0; iter < iterationsCap; iter++ {
		for each := range population {
			fitValues[each] = fitnessFunc(population[each])
		}
		eliteCounter := int(float64(len(population)) * eliteFactor)
		// elites
		elites := eliteSelectorFunc(population, eliteCounter, fitValues)
		// roulette temp population
		tempPopulation := rouletteSelectorFunc(population, len(population), fitValues)

		offspring := make([]Chromosome, len(tempPopulation))
		for i := 0; i < len(population); i++ {
			offspring[i] = crossingFunc(tempPopulation[i], tempPopulation[len(offspring)-1-i])[0]
		}

		offspring = sortChromosomesByFitDesc(offspring, fitValues)
		for each := range elites {
			offspring[len(offspring)-1-each] = elites[each]
		}

		population = deepCopy(offspring)
		for each := range population {
			population[each] = mutationFunc(population[each], mutationProb)
		}
	}

	return maxFromPopulation(population, fitnessFunc)
}

func eachGenMutation(chromosome Chromosome, prob float64) Chromosome {
	for each := range chromosome {
		probVal := rand.Float64()
		if probVal < prob {
			if chromosome[each] == 0.0 {
				chromosome[each] = 1.0
			} else {
				chromosome[each] = 0.0
			}
		}
	}

	return chromosome
}

func travellerGenerator(chromLen int) Chromosome {
	rand.Seed(time.Now().UnixNano())
	newborn := make([]float64, chromLen)

	helper := make([]float64, chromLen)
	for each := range helper {
		helper[each] = float64(each)
	}
	for each := range newborn {
		random := rand.Intn(len(helper))
		newborn[each] = helper[random]
		helper = append(helper[:random], helper[random+1:]...)
	}

	return newborn
}

// distance
func travellerProblemFitness(chromosome Chromosome) float64  {
	coords := []Pair{
		{119, 38},
		{37, 38},
		{197, 55},
		{85, 165},
		{12, 50},
		{100, 53},
		{81, 142},
		{121, 137},
		{85, 145},
		{80, 197},
		{91, 176},
		{106, 55},
		{123, 57},
		{40, 81},
		{78, 125},
		{190, 46},
		{187, 40},
		{37, 107},
		{17, 11},
		{67, 56},
		{78, 133},
		{87, 23},
		{184, 197},
		{111, 12},
		{66,178},
	}
	prev := coords[0]
	sum := 0.0
	ind := 0
	for each := range chromosome {
		ind = int(chromosome[each])
		sum += math.Sqrt(math.Pow(prev.x - coords[ind].x, 2) + math.Pow(prev.y - coords[ind].y, 2))
		prev = coords[ind]
	}
	sum += math.Sqrt(math.Pow(prev.x - coords[0].x, 2) + math.Pow(prev.y - coords[0].y, 2))

	return 1 / sum
}

func swapNeighbourMutation(chromosome Chromosome, prob float64) Chromosome {
	for i := 0; i < len(chromosome) - 1; i++ {
		random := rand.Float64()
		if random < prob {
			temp := chromosome[i]
			chromosome[i] = chromosome[i + 1]
			chromosome[i + 1] = temp
		}
	}
	return chromosome
}

func orderCrossing(mother Chromosome, father Chromosome) []Chromosome {
	child := make(Chromosome, 0)
	firstPoint := rand.Intn(len(mother))
	secondPoint := rand.Intn(len(mother) - firstPoint) + firstPoint
	motherPart := mother[firstPoint:secondPoint]
	f := 0
	for c:= 0; c < firstPoint; f++ {
		if !isIn(motherPart, father[f]) {
			child = append(child, father[f])
			c++
		}
	}

	child = append(child, motherPart...)

	for ; f < len(father); f++{
		if !isIn(child, father[f]) {
			child = append(child, father[f])
		}
	}

	return []Chromosome{child}
}

func isIn(chromosome Chromosome, val float64) bool {
	for each := range chromosome {
		if chromosome[each] == val {
			return true
		}
	}
	return false
}