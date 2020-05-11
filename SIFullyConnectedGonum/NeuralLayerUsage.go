package SIFullyConnectedGonum

import "gonum.org/v1/gonum/mat"

func (layer *NeuralLayer) predict(input *mat.Dense) *mat.Dense {
	output := &mat.Dense{}
	output.Mul(input, layer.Neurons.T())

	x, _ := output.Dims()
	for i := 0; i < x; i++ {
		temp1 := mat.VecDenseCopyOf(output.RowView(i)).RawVector().Data
		temp := layer.ActiveFunc(temp1)
		output.SetRow(i, temp)
	}

	return output
}

func (layer *NeuralLayer) fit(alpha float64, input *mat.Dense,  deltas *mat.Dense) {
	deltas.Mul(deltas, input) // Transpose (T()) might be necessary
	deltas.Scale(alpha, deltas)
	layer.Neurons.Sub(layer.Neurons, deltas)
}