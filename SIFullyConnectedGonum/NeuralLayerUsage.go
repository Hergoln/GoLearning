package SIFullyConnectedGonum

import "gonum.org/v1/gonum/mat"

func (layer *NeuralLayer) predict(input *mat.Dense) *mat.Dense {
	output := &mat.Dense{}
	output.Mul(input, layer.Neurons.T())

	x, _ := output.Dims()
	for i := 0; i < x; i++ {
		temp := mat.VecDenseCopyOf(output.RowView(i)).RawVector().Data
		if layer.ActiveFunc != nil {
			temp = layer.ActiveFunc(temp)
		}
		output.SetRow(i, temp)
	}

	return output
}

func (layer *NeuralLayer) fit(alpha float64, input *mat.Dense, deltas *mat.Dense) {
	tempu := &mat.Dense{}
	tempu.Mul(deltas.T(), input)
	tempu.Scale(alpha, tempu)
	layer.Neurons.Sub(layer.Neurons, tempu)
}
