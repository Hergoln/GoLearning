package SIFullyConnected

import (
	Matrix "../GoMatrixUtils"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	layersIO := input
	for i := 0; i < len(network.Layers) - 1; i++ {
		layersIO = network.Layers[i].predict(input)
	}
	return network.Layers[len(network.Layers) - 1].scale(layersIO)
}

func (network *DeepNeuralNet) Fit(alpha float64, goals []float64, input []float64) float64 {
	// outputs
	lensCount := len(network.Layers)
	layersSizes := make([]int, lensCount)
	for layer := range layersSizes {
		layersSizes[layer] = len(network.Layers[layer].Neurons)
	}
	outputs := network.outputs(input)
	outputs = dropout(outputs) // <- maybe create something like dropout strategy or something?

	deltas  := make([][][]float64, len(network.Layers))
	deltasDeriv := make([][]float64, len(network.Layers))

	// deltas
	deltas[lensCount - 1] = network.outDelta(outputs[lensCount], goals)
	for i := lensCount - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])
	}

	// error
	errors  := .0
	for neuron := range deltas[lensCount - 1][0] {
		errors += deltas[lensCount - 1][0][neuron] * deltas[lensCount - 1][0][neuron]
	}

	// deltas derivatives
	// for now output layer will be treated separately, output layer delta will be passed
	// as active function derivative parameter instead of this layers input
	// SOMETHING WRONG GOING ON HERE, WITHOUT THIS BIT IT WORKS FINE
	// outlayer derivative
	if network.Layers[lensCount - 1].DerivFunc != nil {
		deltasDeriv[lensCount - 1] =  network.Layers[lensCount - 1].DerivFunc(outputs[lensCount])
		//deltasDeriv[lensCount - 1] =  network.Layers[lensCount - 1].DerivFunc(outputs[lensCount])
	} else {
		deltasDeriv[lensCount - 1] = Matrix.OnesVector(len(deltas[lensCount - 1][0]))
	}
	for each := range deltasDeriv[lensCount - 1] {
		deltas[lensCount - 1][0][each] *= deltasDeriv[lensCount - 1][each]
	}

	// hidden layers derivatives
	for i := lensCount - 1; i > 0; i-- {
		if network.Layers[i - 1].DerivFunc != nil {
			deltasDeriv[i - 1] =  network.Layers[i - 1].DerivFunc(outputs[i])
			for each := range deltas[i - 1][0] {
				deltas[i - 1][0][each] *= deltasDeriv[i-1][each]
			}
		}
	}

	// fitting
	for i := range network.Layers {
		network.Layers[i].fit(alpha, outputs[i], deltas[i])
	}
	return errors
}

func (network *DeepNeuralNet) outputs(input []float64) [][]float64 {
	layersCount := len(network.Layers)
	outputs := make([][]float64, layersCount + 1)
	outputs[0] = input
	for i := 0; i < layersCount; i++ {
		outputs[i+1] = network.Layers[i].predict(outputs[i])
	}
	return outputs
}

func (network *DeepNeuralNet) outDelta(netOut, netExp []float64) [][]float64{
	layersCount := len(network.Layers)
	deltas := make([][]float64, 1)
	deltas[0] = make([]float64, len(netOut))
	for i := range network.Layers[layersCount - 1].Neurons {
		deltas[0][i] = netOut[i] - netExp[i]
	}
	return deltas
}

func hidDelta(nextLayer NeuralLayer, nextLayerDelta [][]float64) [][]float64{
	return layerDeltaMul(nextLayer, nextLayerDelta)
}

func layerDeltaMul(layer NeuralLayer, delta [][]float64) [][]float64 {
	result := make([][]float64, 1)
	result[0] = make([]float64, len(layer.Neurons[0].Weights))

	for i := range result[0] {
		for k := range layer.Neurons {
			result[0][i] += delta[0][k] * layer.Neurons[k].Weights[i]
		}
	}

	return result
}

func (network DeepNeuralNet) Description() string {
	var b strings.Builder
	for iL, layer := range network.Layers  {
		fmt.Fprintf(&b, "{ Layer %d\n", iL)
		for iN, neuron := range layer.Neurons {
			fmt.Fprintf(&b, "\tN%d", iN)
			fmt.Fprintln(&b, neuron)
		}
		fmt.Fprintln(&b, "}")
	}
	return b.String()
}

func (network *DeepNeuralNet) DisplayNet() {
	for iL, layer := range network.Layers  {
		fmt.Printf("{ Layer %d\n", iL)
		for iN, neuron := range layer.Neurons {
			fmt.Printf("\tN%d", iN)
			fmt.Println(neuron)
		}
		fmt.Println("}")
	}
}

func dropout(outputs [][]float64) [][]float64  {
	rand.Seed(time.Now().UnixNano())
	counter := 0.0
	randVal := 0.0
	for layer := 1; layer < len(outputs) - 1; layer++ {
		counter = 0.0
		for neuron := range outputs[layer] {
			randVal = float64(rand.Uint32() % 2)
			counter += randVal
			outputs[layer][neuron] *= randVal
		}

		for neuron := range outputs[layer] {
			outputs[layer][neuron] *= float64(len(outputs[layer])) / counter
		}
	}
	return outputs
}

//func dropout(layersLens []int) [][]float64  {
//	rand.Seed(time.Now().UnixNano())
//	onNeuronCounter := 0.0
//	vector := make([][]float64, len(layersLens) - 1) // output layer is not affected by dropout thus -1
//	for i := range vector {
//		vector[i] = make([]float64, layersLens[i])
//		for j := range vector[i] {
//			vector[i][j] = float64(rand.Uint32() %2)
//			onNeuronCounter += vector[i][j]
//		}
//
//		for j := range vector[i] {
//			vector[i][j] *= float64(layersLens[i]) / onNeuronCounter
//		}
//	}
//	return vector
//}
/*
	This version of FitBatch is using my own
	All this have no sense because I have not thought of an idea of treating
	neuron as just a number in matrix and used this freaking encapsulation
	idea, which prevents me from doing matrix computations directly on
	matrices(layers)
 */

func (network *DeepNeuralNet) FitBatch(alpha float64, goals [][]float64, inputs [][]float64) float64 {
	batchSize := len(inputs)
	lensCount := len(network.Layers)
	layersSizes := make([]int, lensCount)
	for layer := range layersSizes {
		layersSizes[layer] = len(network.Layers[layer].Neurons)
	}

	// 1st dim -> batches (outptus[0] = first batch)
	// 2nd dim -> layers outputs (outptus[0][0] = first batch net input;
	//							  outptus[0][1] = first batch, first layer outptut)
	// 3rd dim -> input and outputs to/from neurons (outptus[0][0][0] = input of first weight of first neuron;
	//												 outptus[0][1][0] = output of first neuron
	outputs := make([][][]float64, batchSize)
	for batch := range outputs {
		// this function also handles activation functions step
		outputs[batch] = network.outputs(inputs[batch])
	}

	for batch := range outputs {
		// dropout does not modify first matrix(network input) and last matrix(output layer)
		outputs[batch] = dropout(outputs[batch])
	}

	deltas  := make([][][][]float64, batchSize)
	deltasDeriv := make([][][]float64, batchSize)

	for each := 0; each < batchSize; each++ {
		deltas[each] = make([][][]float64, lensCount)
		deltasDeriv[each] = make([][]float64, lensCount)

		deltas[each][lensCount - 1] = network.outDelta(outputs[each][lensCount], goals[each])
		for i := lensCount - 1; i > 0; i-- {
			deltas[each][i-1] = hidDelta(network.Layers[i], deltas[each][i])
		}
	}

	errNet  := .0
	for batch := range outputs {
		for neuron := range deltas[lensCount - 1][0] {
			errNet += deltas[batch][lensCount - 1][0][neuron] * deltas[batch][lensCount - 1][0][neuron]
		}
	}

	for batch := range outputs {
		if network.Layers[lensCount - 1].DerivFunc != nil {
			deltasDeriv[batch][lensCount - 1] =  network.Layers[lensCount - 1].DerivFunc(outputs[batch][lensCount])
			//deltasDeriv[lensCount - 1] =  network.Layers[lensCount - 1].DerivFunc(outputs[lensCount])
		} else {
			deltasDeriv[batch][lensCount - 1] = Matrix.OnesVector(len(deltas[batch][lensCount - 1][0]))
		}
		for each := range deltasDeriv[lensCount - 1] {
			deltas[batch][lensCount - 1][0][each] *= deltasDeriv[batch][lensCount - 1][each]
		}

		// hidden layers derivatives
		for i := lensCount - 1; i > 0; i-- {
			if network.Layers[i - 1].DerivFunc != nil {
				deltasDeriv[batch][i - 1] =  network.Layers[i - 1].DerivFunc(outputs[batch][i])
				for each := range deltas[i - 1][0] {
					deltas[batch][i - 1][0][each] *= deltasDeriv[batch][i-1][each]
				}
			}
		}
	}

	for batch := range outputs {
		for i := range network.Layers {
			network.Layers[i].fit(alpha, outputs[batch][i], deltas[batch][i])
		}
	}


	return errNet
}

func PredictBatch(input [][]float64) {
	panic("NotImplemented")
}

func (network DeepNeuralNet) CheckError(expectedOutputs, input []float64) {
	panic("NotImplemented")
}