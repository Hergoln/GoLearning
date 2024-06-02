package ConvolutionalTests

import (
	fun "../FunctionsAndDerivatives"
	Conv "../SIConvolutional"
	SIG "../SIFullyConnectedGonum"
	"gonum.org/v1/gonum/mat"
)

func Example() {
	//input := []float64 {
	//	8.5, 0.65, 1.2,
	//	9.5, 0.8, 1.3,
	//	9.9, 0.8, 0.5,
	//	9.0, 0.9, 1.0,
	//}

	input := []float64 {
		8.5, 9.5, 9.9, 9.0,
		0.65, 0.8, 0.8, 0.9,
		1.2, 1.3, 0.5, 1.0,
	}

	expected := []float64 {0, 1}

	conv := Conv.CreateConvLayer([][]float64{
		{0.1, 0.2, -0.1, -0.1, 0.1, 0.9, 0.1, 0.4, 0.1},
		{0.3, 1.1, -0.3, 0.1, 0.2, 0.0, 0.0, 1.3, 0.1},
	})

	fc := createLayer(
		[][]float64{
			{0.1, -0.2, 0.1, 0.3},
			{0.2, 0.1, 0.5, -0.3},
		})

	Conv.ConvAndFcFit(
		0.01,
		fc,
		conv,
		input,
		3,
		4,
		expected,
		fun.ReLu,
		fun.ReLuDeriv,
	)
}

func createLayer(weights [][]float64) *SIG.NeuralLayer {
	data := make([]float64, 0)
	for each := range weights {
		data = append(data, weights[each]...)
	}
	newborn := &SIG.NeuralLayer{
		Neurons:   mat.NewDense(len(weights), len(weights[0]), data),
		ActiveFunc: nil,
		DerivFunc:  nil,
	}
	return newborn
}