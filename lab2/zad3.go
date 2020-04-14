package lab2

import (
	"../LegacySIFullyConnected"
	"fmt"
)

func Zad3() {
	layer := LegacySIFullyConnected.GibNeurons(
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

	alpha := 0.01

	//fmt.Println(network.Layers[0].Study(alpha, expectedSeries[0], inputSeries[0]))
	//for i := range network.Layers[0].Neurons {
	//	fmt.Println(network.Layers[0].Neurons[i])
	//}

	err := .0
	for i := 0; i < 80000; i++ {
		var errors []float64
		err = .0
		for j := range expectedSeries {
			errors = layer.Study(alpha, expectedSeries[j], inputSeries[j])
			for _, e := range errors {
				err += e
			}
		}
	}
	fmt.Println(err)

	/*
		Q: Jaka jest minimalna wartość błędu skumulowanego, jaką może osiągnąć ta sieć?
		A: 0.1082564980797942, gdzieś pomiędzy 7 a 8 tysięcy iteracji przy alpha = 0.01
	 */

}