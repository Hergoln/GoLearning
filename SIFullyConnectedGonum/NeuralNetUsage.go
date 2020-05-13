package SIFullyConnectedGonum

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	denseInput := mat.NewDense(1, len(input), input)
	for i := range network.Layers {
		denseInput = network.Layers[i].Predict(denseInput)
	}
	return denseInput.RawRowView(0)
}

func (network DeepNeuralNet) outputs(input []float64) []*mat.Dense {
	layersCount := len(network.Layers)
	outputs := make([]*mat.Dense, layersCount+1)
	outputs[0] = mat.NewDense(1, len(input), input)
	for i := 0; i < layersCount; i++ {
		outputs[i+1] = network.Layers[i].Predict(outputs[i])
	}
	return outputs
}

func (network *DeepNeuralNet) outDelta(netOut, netExp *mat.Dense) *mat.Dense {
	delta := &mat.Dense{}
	delta.Sub(netOut, netExp)
	return delta
}

func hidDelta(nextLayer NeuralLayer, nextLayerDelta *mat.Dense) *mat.Dense {
	result := &mat.Dense{}
	result.Mul(nextLayerDelta, nextLayer.Neurons)
	return result
}

func (network *DeepNeuralNet) Fit(alpha float64, goals []float64, input []float64) float64 {
	layersCount := len(network.Layers)

	outputs := network.outputs(input)
	outputs = network.dropout(outputs)

	deltas := make([]*mat.Dense, layersCount)
	deltasDeriv := make([]*mat.Dense, layersCount)

	// deltas
	deltas[layersCount-1] = network.outDelta(outputs[layersCount], mat.NewDense(1, len(goals), goals))
	for i := layersCount - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])
	}

	//deltas derivatives
	for layer := 0; layer < layersCount; layer++ {
		x, y := deltas[layer].Dims()
		deltasDeriv[layer] = mat.NewDense(x, y, nil)
		for i := 0; i < x; i++ {
			// here might be problem with dims as well
			deltasDeriv[layer].SetRow(i, network.Layers[layer].DerivFunc(mat.VecDenseCopyOf(outputs[layer+1].RowView(i)).RawVector().Data))
		}
	}

	// error
	errors := 0.0
	x, _ := deltas[layersCount-1].Dims()
	for i := 0; i < x; i++ {
		rawLastDelta := deltas[layersCount-1].RawRowView(i)
		for neuron := range rawLastDelta {
			errors += rawLastDelta[neuron] * rawLastDelta[neuron]
		}
	}

	for each := range deltas {
		deltas[each].MulElem(deltas[each], deltasDeriv[each])
	}

	// fitting
	for i := range network.Layers {
		network.Layers[i].fit(alpha, outputs[i], deltas[i])
	}
	return errors
}

func (network *DeepNeuralNet) dropout(outputs []*mat.Dense) []*mat.Dense {
	rand.Seed(time.Now().UnixNano())
	counter := 0.0
	randVal := 0.0
	var mask []float64
	var maskDense *mat.Dense
	// dropout does not modify first matrix(network input) and last matrix(output layer)
	for layer := 1; layer < len(outputs)-1; layer++ {
		counter = 0.0
		x, y := outputs[layer].Dims()
		mask = make([]float64, x*y)
		for i := range mask {
			randVal = network.DropoutStrategy()
			counter += randVal
			mask[i] = randVal
		}
		maskDense = mat.NewDense(x, y, mask)
		maskDense.Scale(float64(len(mask))/counter, maskDense)
		outputs[layer].MulElem(maskDense, outputs[layer])
	}
	return outputs
}

func (network *DeepNeuralNet) FitBatch(alpha float64, goals [][]float64, inputs [][]float64) float64 {
	layersCount := len(network.Layers)

	outputs := network.outputsBatch(inputs)
	outputs = network.dropout(outputs)

	deltas := network.calcDeltas(goals, outputs[layersCount], layersCount)
	deltasDeriv := network.calcDeltasDeriv(deltas, outputs, layersCount)

	for each := range deltas {
		deltas[each].MulElem(deltas[each], deltasDeriv[each])
	}

	for i := range network.Layers {
		network.Layers[i].fit(alpha, outputs[i], deltas[i])
	}

	errors := 0.0
	x, _ := deltas[layersCount-1].Dims()
	for i := 0; i < x; i++ {
		rawLastDelta := deltas[layersCount-1].RawRowView(i)
		for neuron := range rawLastDelta {
			errors += rawLastDelta[neuron] * rawLastDelta[neuron]
		}
	}

	return errors
}

func (network DeepNeuralNet) outputsBatch(inputs [][]float64) []*mat.Dense {
	layersCount := len(network.Layers)
	outputs := make([]*mat.Dense, layersCount+1)
	var vector []float64

	for row := range inputs {
		vector = append(vector, inputs[:][row]...)
	}

	outputs[0] = mat.NewDense(len(inputs), len(inputs[0]), vector)
	for i := 0; i < layersCount; i++ {
		outputs[i+1] = network.Layers[i].Predict(outputs[i])
	}
	return outputs
}

func (network *DeepNeuralNet) calcDeltas(goals [][]float64, outputs *mat.Dense, layersCount int) []*mat.Dense {

	deltas := make([]*mat.Dense, layersCount)

	var vector []float64
	for row := range goals {
		vector = append(vector, goals[:][row]...)
	}
	temp := mat.NewDense(len(goals), len(goals[0]), vector)
	deltas[layersCount-1] = network.outDelta(outputs, temp)
	for i := layersCount - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])
	}

	return deltas
}

func (network *DeepNeuralNet) calcDeltasDeriv(deltas []*mat.Dense, outputs []*mat.Dense, layersCount int) []*mat.Dense {

	deltasDeriv := make([]*mat.Dense, layersCount)

	for layer := 0; layer < layersCount; layer++ {
		x, y := deltas[layer].Dims()
		deltasDeriv[layer] = mat.NewDense(x, y, nil)
		for i := 0; i < x; i++ {
			// here might be problem with dims as well
			deltasDeriv[layer].SetRow(i, network.Layers[layer].DerivFunc(mat.VecDenseCopyOf(outputs[layer+1].RowView(i)).RawVector().Data))
		}
	}

	return deltasDeriv
}
