package ConvolutionalTests

import (
	fun "../FunctionsAndDerivatives"
	MNIST "../MnistDBUtils"
	CONV "../SIConvolutional"
	FC "../SIFullyConnectedGonum"
	"fmt"
	"math/rand"
)

func Zad2() {
	// don't know how to get this value, thus its hard coded for now
	// network creation
	alpha := 0.02
	inputSize := 676 * 16
	inputRows := 28
	inputCols := 28
	outputSize := 10

	trainLabels := MNIST.ParseLabelFile("/train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("/train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("/t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("/t10k-images.idx3-ubyte")

	fc := FC.CreateNetwork(
		[]int{inputSize, outputSize},
		[]FC.ActiveFunc{nil},
		[]FC.ActiveFunc{nil},
		func() float64 { return rand.Float64()*0.02 - 0.009 },
		func() float64 { return float64(rand.Uint64() % 2) },
	)

	conv := CONV.RandConvLayer(3, 3, 16, func() float64 { return rand.Float64()*0.02 - 0.01 })

	netErr := 0.0
	for i, limit := 0, 100; i < limit; i++ {
		netErr = 0.0
		for set := range trainLabels.Labels {
			netErr = CONV.ConvAndFcFit(
				alpha,
				&fc.Layers[0],
				conv,
				MNIST.GetInputVector(trainImages.Images[set]),
				inputRows,
				inputCols,
				MNIST.GetExpectedVector(trainLabels.Labels[set]),
				fun.ReLu,
				fun.ReLuDeriv,
			)
		}

		var prediction []float64
		errorCounter := 0
		for set := range testLabels.Labels {
			prediction = CONV.ConvAndFcPredict(
				&fc.Layers[0],
				conv,
				MNIST.GetInputVector(testImages.Images[set]),
				inputRows,
				inputCols,
			)
			if testLabels.Labels[set] != MNIST.GetOutputLabel(prediction) {
				errorCounter++
			}
		}

		if i%5 == 0 {
			fmt.Printf(
				"%d iteration, network error: %f network score: %d/%d\n",
				i,
				netErr,
				int(testLabels.DataCount)-errorCounter,
				testLabels.DataCount,
			)
		}

	}

}
