package SIConvolutional

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func Pooling(maskSize, stride int) {
	panic("NotImplemented")
}

func Convolute(input *mat.Dense, filters *mat.Dense, stride, padding int) (*mat.Dense, int, int) {
	_, filterCols := filters.Dims()
	filterSize := int(math.Sqrt(float64(filterCols)))
	imageSections, sectionRows, sectionCols := extractImageSections(input, filterSize, stride)

	// no sizes might cause problems
	output := &mat.Dense{}
	output.Mul(imageSections, filters.T())
	return output, sectionRows, sectionCols
}

// I am assuming all filters will be squares, it reduces workload
func extractImageSections(input *mat.Dense, filterSize, stride int) (*mat.Dense, int, int) {
	imageRawData := input.RawMatrix().Data
	inRows, inCols := input.Dims()
	//* W2 = (W1 − F)/S + 1,
	//• H2 = (H1 − F)/S + 1,
	//• D2 = D1
	outRowsCount := (inRows - filterSize)/stride + 1
	outColsCount := (inCols - filterSize)/stride + 1

	outputData := make([]float64, 0, filterSize * filterSize * outRowsCount * outColsCount)
	for j := 0; j < outColsCount; j += stride {
		for i := 0; i < outRowsCount; i += stride{
			outputData = append(outputData, extractMask(imageRawData, filterSize, inCols, i, j)...)
		}
	}

	output := mat.NewDense(outRowsCount * outColsCount, filterSize * filterSize, outputData)
	return output, outRowsCount, outColsCount
}

// data -> flatten image
func extractMask(data []float64, maskSize, rowLen, offsetRow, offsetCol int) []float64 {
	out := make([]float64, 0, maskSize * maskSize)

	for i := 0; i < maskSize; i++ {
		beg := offsetRow + rowLen * (offsetCol + i)
		out = append(out, data[beg : beg + maskSize]...)
	}

	return out
}

func Describe(data *mat.Dense) string {
	rows, _ := data.Dims()
	output := ""

	for i := 0; i < rows; i++ {
		for _, j := range data.RawRowView(i){
			output += fmt.Sprintf("%.2f ", j)
		}
		output += "\n"
	}

	return output
}