package SI

import (
	"math"
	"math/rand"
)

type INeuralNet interface {
	Predict(input []float64) float64
	Study(goals float64, inputs []float64)
	Dropout()
	PredictBatch(packages interface{})
}

// lab 2: zad 3,4
type NeuralLayer struct {
	Neurons []Neuron
}

func (layer NeuralLayer) Predict(input []float64) []float64 {
	output := make([]float64, len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		output[i] = neuron.Scale(input)
	}
	return output
}

func (layer *NeuralLayer) Study(alpha float64, goals []float64, inputs []float64) []float64 {
	errors := make([]float64, len(goals))
	for i := range layer.Neurons {
		errors[i] = layer.Neurons[i].Study(alpha, goals[i], inputs)
	}
	return errors
}

// lab 3
type DeepNeuralNet struct {
	Layers []NeuralLayer
}

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	layersIO := input
	for _, layer := range network.Layers {
		layersIO = layer.Predict(layersIO)
	}
	return layersIO
}

func ReLu(x float64) float64 {
	return math.Min(0, x)
}

func (network *DeepNeuralNet) Study(goals, inputs []float64) []float64 {
	panic("NotImplemented")
}

func (network DeepNeuralNet) Dropout() {
	panic("NotImplemented")
}

func (network DeepNeuralNet) PredictBatch(packages interface{}) {
	panic("NotImplemented")
}

func GibWeights(weights []float64) Neuron {
	born := Neuron{}
	for _, weight := range weights {
		born.Weights = append(born.Weights, weight)
	}
	return born
}

func GibNeurons(weights [][]float64) NeuralLayer {
	born := NeuralLayer{}
	for i := range weights {
		born.Neurons = append(born.Neurons, GibWeights(weights[i]))
	}
	return born
}

func GibLayers(weights [][][]float64) DeepNeuralNet {
	born := DeepNeuralNet{}
	for i := range weights {
		born.Layers = append(born.Layers, GibNeurons(weights[i]))
	}
	return born
}

func ConstructRandomNetwork(inputsCount, outputsCount int) DeepNeuralNet {
	firstLayer := make([]Neuron, outputsCount)
	for i := range firstLayer {
		firstLayer[i] = Neuron{}
		dendrites := make([]float64, inputsCount)
		for j := range dendrites {
			dendrites[j] = rand.Float64()
		}
		firstLayer[i].Weights = dendrites
	}
	return DeepNeuralNet{[]NeuralLayer{NeuralLayer{firstLayer}}}
}

func (network *DeepNeuralNet) AppendRandomLayer(outputNeuronsNumber int) {
	neurons := make([]Neuron, outputNeuronsNumber)
	newLayer := NeuralLayer{neurons}
	connectionsNumber := len(network.Layers[len(network.Layers)-1].Neurons)

	for i := range neurons {
		dendrites := make([]float64, connectionsNumber)
		for j := range dendrites {
			dendrites[j] = rand.Float64()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func (network DeepNeuralNet) CheckError(expectedOutputs, inputs []float64) {
	panic("NotImplemented")
}