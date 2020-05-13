package ConvolutionalTests

import (
	"../SIConvolutional"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func Zad1() {

	inputImage := []float64{
		1., 1., 1., 0., 0.,
		0., 1., 1., 1., 0.,
		0., 0., 1., 1., 1.,
		0., 0., 1., 1., 0.,
		0., 1., 1., 0., 0.,
	}

	mask := []float64{
		1., 0., 1.,
		0., 1., 0.,
		1., 0., 1.,
	}

	filters := mat.NewDense(1, len(mask), mask)

	output, secRows, secCols := SIConvolutional.Convolute(
		mat.NewDense(5, 5, inputImage),
		filters,
		1,
		0,
	)

	output = mat.NewDense(secRows, secCols, output.RawMatrix().Data)

	fmt.Println(SIConvolutional.Describe(output))

}
