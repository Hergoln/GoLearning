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
	//lab4Zad1()
	lab4Zad2_3()
}

func lab4Zad2_3() {
	rand.Seed(time.Now().UnixNano())
	trainLabels := MNIST.ParseLabelFile("train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("t10k-images.idx3-ubyte")

	alpha := 0.0001
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
	for j, limit := 0, 100; j < limit; j++ {
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

		if j%5 == 0 {
			fmt.Printf(
				"%d iteration, network error: %f network score: %d/%d\n",
				j,
				netErr,
				int(testLabels.DataCount)-errorCounter,
				testLabels.DataCount,
			)
		}
	}
}

func lab4Zad1() {
	rand.Seed(time.Now().UnixNano())
	trainLabels := MNIST.ParseLabelFile("train-labels.idx1-ubyte")
	trainImages := MNIST.ParseImageFile("train-images.idx3-ubyte")

	testLabels := MNIST.ParseLabelFile("t10k-labels.idx1-ubyte")
	testImages := MNIST.ParseImageFile("t10k-images.idx3-ubyte")

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

	alpha := 0.005
	netErr := 0.0
	for i, limit := 0, 350; i < limit; i++ {
		netErr = 0.0
		for i := range trainLabels.Labels {
			netErr = network.Fit(
				alpha,
				MNIST.GetExpectedVector(trainLabels.Labels[i]),
				MNIST.GetInputVector(trainImages.Images[i]),
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