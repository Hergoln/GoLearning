package SIFullyConnectedGonum

import "gonum.org/v1/gonum/mat"

func (layer *NeuralLayer) Predict(input *mat.Dense) *mat.Dense {
	output := &mat.Dense{}
	output.Mul(input, layer.Neurons.T())

	if layer.ActiveFunc != nil {
		x, _ := output.Dims()
		for i := 0; i < x; i++ {
			temp := mat.VecDenseCopyOf(output.RowView(i)).RawVector().Data
			temp = layer.ActiveFunc(temp)
			output.SetRow(i, temp)
		}
	}

	return output
}

func (layer *NeuralLayer) fit(alpha float64, input *mat.Dense, deltas *mat.Dense) {
	tempu := &mat.Dense{}
	tempu.Mul(deltas.T(), input)
	tempu.Scale(alpha, tempu)
	layer.Neurons.Sub(layer.Neurons, tempu)
}
