package lab3

import (
	"../LegacySIFullyConnected"
	"fmt"
	"math"
)

/*
	należy zapamiętywać dla każdej warstwy wyniki a potem delty dla każdej warstwy
	i potem dopiero po wyjściu danych z sieci aktualizujemy całość na raz,
	więc wyjśćie jednej wchodzi w drugą ale te wyjścia też muszą być gdzieś zapisane

	no i na tym polega propagacja wsteczna, po obliczeniu wszystkich warstw i wag
	obliczamy wagi dla wyjśia, potem dla warstwy n-1 na podstawie wyjścia, potem
	dla n-2 na podstawie n-1 itd

	// 1 krok - wyniki dla poszczególnych warstw
	obliczanie dla hidden
	obliczanie dla output

	// 2 krok - delta dla warstw
	output_l_delta = (output - expected) * hidden * alpha
	hidden_l_delta = output_l_weights * (output - expected) * alpha

	// 3 krok - aktualizacja
	output_l_weights -= output_l_delta * layer_input
	hidden_l_weights -= hidden_l_delta * layer_input


	output_l_weights -= hidden_output 	 * (output - expected)	* alpha * layer_input
	hidden_l_weights -= output_l_weights * (output - expected) 	* alpha * layer_input

	lab3 zad1 wyjście musi jeszcze przejść przez funkcję aktywacji

	lab3 zad4 obrazem przekonwertować jako wektor, wyjśćie określa z jaką pewnością
	dany obrazek jest daną cyfrą (10 neuronów)

	wagi mają być od -0.1 do 0.1
	alpha = 0.01
	dane należy przeskalować do wartości od 0 - 1

	pozostawienie wartości od 0 - 255 sprawi, że wagi będą bardzo szybko rosły i dłuugo
	się będzie uczyło i będzie bardzo niestabilna

 */

func Zad1_2() {
	network := LegacySIFullyConnected.GibLayers(
		[][][]float64{
			{
				{0.1, 0.2, -0.1},
				{-0.1, 0.1, 0.9},
				{0.1, 0.4, 0.1},
			},
			{
				{0.3, 1.1, -0.3},
				{0.1, 0.2, 0.0},
				{0.0, 1.3, 0.1},
			},
		})

	inputSeries := [][]float64{
		{8.5, 0.65, 1.2},	// serie 1
		{9.5, 0.8, 1.3},	// serie 2
		{9.9, 0.8, 0.5},	// serie 3
		{9.0, 0.9, 1.0},	// serie 4
	}

	for i := range inputSeries {
		fmt.Printf("%.3f\n", network.PredictActiveFunc(inputSeries[i], ReLu))
	}

	expectedSeries := [][]float64{
		{0.1, 1.0, 0.1},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 0.1},
		{0.1, 1.0, 0.2},
	}

	alpha := 0.01

	//network.StudyActiveFunc(alpha, expectedSeries[0], inputSeries[0], ReLu)
	err := 0.0
	for i := 0; i <= 50; i++ {
		err = 0.0
		for series := range expectedSeries {
			temp := network.StudyActiveFunc(alpha, expectedSeries[series], inputSeries[series], ReLu)
			fmt.Printf("%f +\n", temp)
			err += temp
		}
	}

	fmt.Printf("end, cumulative err: %f", err)
}

func ReLu(x float64) float64 {
	return math.Max(0, x)
}