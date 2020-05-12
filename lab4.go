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
	trainLabels := MNIST.ParseLabelFile("train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("t10k-images.idx3-ubyte")

	alpha := 0.0005
	batchSize := 100
	inputSize := int(trainImages.Width * trainImages.Height)
	hiddenLayer := 100
	outputSize := 10

	network := SI.CreateNetwork(
		//inputLayer, hiddenLayer, outputLayer
		[]int{inputSize, hiddenLayer, outputSize},
		[]SI.ActiveFunc{fun.Tanh, fun.Softmax},
		[]SI.ActiveFunc{fun.TanhDeriv, fun.SoftmaxDeriv},
		func() float64 { return rand.Float64()*0.02 - 0.01 },
		func() float64 { return float64(rand.Uint64() % 2) },
	)

	netErr := 0.0
	for i, limit := 0, 100; i < limit; i++ {
		netErr = 0.0
		for i := 0; i < len(trainLabels.Labels); i += batchSize {
			batchGoals := trainLabels.Labels[i : i+batchSize]
			batchInputs := trainImages.Images[i : i+batchSize]

			netErr = network.FitBatch(
				alpha,
				MNIST.GetExpectedMatrix(batchGoals),
				MNIST.GetInputMatrix(batchInputs),
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
