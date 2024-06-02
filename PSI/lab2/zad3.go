package lab2

import (
	"../LegacySIFullyConnected"
	"fmt"
)

func Zad3() {
	// 2.3
	layerHuan := LegacySIFullyConnected.GibNeurons(
		[][]float64{
			{0.1, 0.1, -0.3},
			{0.1, 0.2, 0.0},
			{0.0, 1.3, 0.1},
		})

	layerTwuan := LegacySIFullyConnected.GibNeurons(
		[][]float64{
			{0.1, 0.1, -0.3},
			{0.1, 0.2, 0.0},
			{0.0, 1.3, 0.1},
		})

	inputSeries := [][]float64{
		{8.5, 0.65, 1.2},	// serie 1
		{9.5, 0.8, 1.3},	// serie 2
		{9.9, 0.8, 0.5},	// serie 3
		{9.0, 0.9, 1.0},	// serie 4
	}

	expectedSeries := [][]float64{
		{0.1, 1.0, 0.1},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 0.1},
		{0.1, 1.0, 0.2},
	}

	fmt.Print("\n2.3\n")
	iterations := 8000
	alpha := 0.01
	studyAndPrint(alpha, iterations, layerHuan, expectedSeries, inputSeries)

	// learning with lower alpha
	alpha = 0.001
	studyAndPrint(alpha, iterations, layerTwuan, expectedSeries, inputSeries)

	/*
		Q: Jaka jest minimalna wartość błędu skumulowanego, jaką może osiągnąć ta sieć?
		A: 0.106491, gdzieś pomiędzy 7 a 8 tysięcy iteracji przy alpha = 0.01
			przy alfie 0.001 bląd może wynieść ok 0.082125
	 */

}

func studyAndPrint(alpha float64, iterations int, layer LegacySIFullyConnected.NeuralLayer, expected [][]float64, input [][]float64) {
	minErr := 1.0
	err := 0.0
	for i := 0; i < iterations; i++ {
		var errors []float64
		err = .0
		for j := range expected {
			errors = layer.Study(alpha, expected[j], input[j])
			for _, e := range errors {
				err += e
			}
		}
		if minErr > err {
			minErr = err
		}
	}

	fmt.Printf("alpha = %.3f\n err = %f\n", alpha, minErr)
}