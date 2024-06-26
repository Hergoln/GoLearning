package main

import (
	fun "./FunctionsAndDerivatives"
	MNIST "./MnistDBUtils"
	SI "./SIFullyConnectedGonum"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// network creation
	alpha := 0.01
	network := SI.CreateNetwork(
		[]int{784, 40, 10},
		[]SI.ActiveFunc{fun.Sigmoid, fun.Softmax},
		[]SI.ActiveFunc{fun.SigmoidDeriv, fun.SoftmaxDeriv},
		func() float64 { return rand.NormFloat64()*0.2 - 0.1 },
		func() float64 { return float64(rand.Uint64() % 2) },
	)

	trainLabels := MNIST.ParseLabelFile("/train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("/train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("/t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("/t10k-images.idx3-ubyte")

	netErr := 0.0
	for i, limit := 0, 100; i < limit; i++ {
		netErr = 0.0
		for set := range trainLabels.Labels {
			netErr = network.Fit(
				alpha,
				MNIST.GetExpectedVector(trainLabels.Labels[set]),
				MNIST.GetInputVector(trainImages.Images[set]),
			)
		}

		var prediction []float64
		errorCounter := 0
		for set := range testLabels.Labels {
			prediction = network.Predict(MNIST.GetInputVector(testImages.Images[set]))
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
