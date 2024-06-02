package SIConvolutional

import(
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func CreateConvLayer(filters [][]float64) *ConvolutionLayer {
	data := make([]float64, 0)

	for each := range filters {
		data = append(data, filters[each]...)
	}
	
	newborn := &ConvolutionLayer{
		Filters: mat.NewDense(len(filters), len(filters[0]), data),
		Type:    0,
	}

	return newborn
}

func RandConvLayer(rows, cols, count int, randStrategy func () float64) *ConvolutionLayer {
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, rows * cols * count)

	for each := range data {
		data[each] = randStrategy()
	}

	newborn := &ConvolutionLayer{
		Filters: mat.NewDense(count, rows * cols, data),
		Type:    0,
	}

	return newborn
}