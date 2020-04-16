package lab3

import (
	fun "../FunctionsAndDerivatives"
	MNIST "../MnistDBUtils"
	SI "../SIFullyConnected"
	"fmt"
	"math/rand"
	"time"
)

func Zad4() {
	// don't know how to get this value, thus its hard coded for now
	rand.Seed(time.Now().UnixNano())
	// network creation
	alpha := 0.01
	network := SI.CreateNetwork(
		[]int{784, 40, 10},
		[]SI.ActiveFunc{fun.Sigmoid},
		[]SI.ActiveFunc{fun.SigmoidDeriv},
		func() float64 {return rand.Float64() * 0.2 - 0.1},
		)

	trainLabels := MNIST.ParseLabelFile("/train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("/train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("/t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("/t10k-images.idx3-ubyte")

	netErr := 0.0
	for i, limit := 0, 100; i < limit; i++{
		netErr = 0.0
		for set := range trainLabels.Labels {
			netErr = network.Fit(
				alpha,
				getExpectedVector(trainLabels.Labels[set]),
				getInputVector(trainImages.Images[set]),
				)
		}

		var prediction []float64
		errorCounter := 0
		for set := range testLabels.Labels {
			prediction = network.Predict(getInputVector(testImages.Images[set]))
			if testLabels.Labels[set] != getOutputLabel(prediction) {
				errorCounter++
			}
		}

		if i%10 == 0 {
				fmt.Printf(
				"%d iteration, network error: %f network score: %d/%d\n",
				i,
				netErr,
				int(testLabels.DataCount) - errorCounter,
				testLabels.DataCount,
				)
		}

	}
}

func getExpectedVector(label byte) []float64 {
	expected := make([]float64, 10)
	expected[label] = 1.0
	return expected
}

func getInputVector(image []byte) []float64 {
	converted := make([]float64, len(image))
	for i := range image {
		converted[i] = float64(image[i]) / 255
	}
	return converted
}

func getOutputLabel(prediction []float64) byte {
	var label byte
	label = 0
	for i := range prediction {
		if prediction[label] < prediction[i] {
			label = byte(i)
		}
	}
	return label
}