package lab2

import (
	"../SI"
	"fmt"
)

func Zad1_2() {
	// zad 1
	input := .1
	alpha := 1.
	goal := .8
	neuron := SI.Neuron{Weights: []float64{.5}}

	for i := 0; i < 2000; i++ {
		fmt.Printf("Error(iter: %d): %f\n", i, neuron.Study(alpha, goal, []float64{input}))
	}
	fmt.Printf("Output: %f\n", neuron.Scale([]float64{input}))

	/*
		Q: Po ilu iteracjach sieć jest w stanie zbliżyć się do oczekiwanego wyniku końcowego?
			Jak myślisz. dlaczego tak się dzieje?
		A: Dla p1 (input = 2, alpha = 1) chyba nigdy, po 325 otrzymałem +Inf natomiast
			po 649 natomiast otrzymuję NaN, nie wiem czemu tak się dzieje, wiem tylko, że
			już po drugiej iteracji waga zaczyna skakać pomiędzy ujemnymi i dodatnimi
			wartościami i do tego coraz większymi
			Dla p2 (input = 0.1, alpha 1) po ok. 694 iteracjach, mój pomysł jest taki, że
			duża ilość operacji jest winą małej w stosunku wartości wejściowej w stosunku
			do wartości oczekiwanej (8 razy mniejsza)
	 */

	// zad 2
	neuron = SI.Neuron{Weights: []float64{0.1, 0.2, -0.1}}
	series := [][]float64{
		{8.5, 0.65, 1.2},	// serie 1
		{9.5, 0.8, 1.3},	// serie 2
		{9.9, 0.8, 0.5},	// serie 3
		{9.0, 0.9, 1.0},	// serie 4
	}
	alpha = 0.01
	expectedOutputs := []float64{1., 1., 0, 1.}
	err := 0.
	for i := 0; i < 15000; i++ {
		err = 0
		for i := range series {
			err += neuron.Study(alpha, expectedOutputs[i], series[i])
		}
	}

	fmt.Println(err)
	fmt.Printf("Summrized error: %f", neuron.SumErrorForSeries(expectedOutputs, series))
	/*
		Q: Po ilu iteracjach błąd będzie mniejszy niż 5%?
		A: Przy alpha = 0.01 przy 15 tysiącach juz jest na pewno mniej niż 5%
	 */
}