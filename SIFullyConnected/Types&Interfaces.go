package SIFullyConnected

type INeuron interface {
	Scale(input []float64) float64
	Study(alpha float64, goal float64, inputs []float64) // for series of data
}

type Neuron struct {
	Weights []float64
}

type ActiveFunc func(float64) float64

type NeuralLayer struct {
	Neurons []Neuron
	ActiveFunc ActiveFunc
	DerivFunc ActiveFunc
}

type INeuralNet interface {
	Predict(input []float64) float64
	Study(goals float64, inputs []float64)
	Dropout()
	PredictBatch(packages interface{})
}

type DeepNeuralNet struct {
	Layers []NeuralLayer
	Alpha float64
}