package SIConvolutional

import (
	Fc "../SIFullyConnectedGonum"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// (this code could be decoupled and than automated to make these functions more flexible)

// I am assuming all filters will be squares, it reduces workload
func extractImageSections(input *mat.Dense, filterSize, stride int, inputRows int, inputCols int) (*mat.Dense, int, int) {
	imageRawData := input.RawMatrix().Data
	//* W2 = (W1 − F)/S + 1,
	//• H2 = (H1 − F)/S + 1,
	//• D2 = D1
	outRowsCount := (inputRows-filterSize)/stride + 1
	outColsCount := (inputCols-filterSize)/stride + 1

	outputData := make([]float64, 0)
	for j := 0; j < outColsCount; j += stride {
		for i := 0; i < outRowsCount; i += stride {
			outputData = append(outputData, extractMask(imageRawData, filterSize, inputCols, j, i)...)
		}
	}

	output := mat.NewDense(outRowsCount*outColsCount, filterSize*filterSize, outputData)
	return output, outRowsCount, outColsCount
}

// data -> flatten image
func extractMask(data []float64, maskSize, rowLen, offsetRow, offsetCol int) []float64 {
	out := make([]float64, 0, maskSize*maskSize)

	for i := 0; i < maskSize; i++ {
		beg := offsetRow + rowLen*(offsetCol+i)
		out = append(out, data[beg:beg+maskSize]...)
	}

	return out
}

func Describe(data *mat.Dense) string {
	rows, _ := data.Dims()
	output := ""

	for i := 0; i < rows; i++ {
		for _, j := range data.RawRowView(i) {
			output += fmt.Sprintf("%.2f ", j)
		}
		output += "\n"
	}

	return output
}

func addPadding(img *mat.Dense, padding int) *mat.Dense {
	if padding == 0 {
		return mat.DenseCopyOf(img)
	}

	data := make([]float64, 0)
	rows, cols := img.Dims()
	for i := 0; i < padding; i++ {
		data = append(data, make([]float64, cols+padding*2)...)
	}

	for i := 0; i < rows; i++ {
		data = append(data, make([]float64, padding)...)
		data = append(data, img.RawRowView(i)...)
		data = append(data, make([]float64, padding)...)
	}

	for i := 0; i < padding; i++ {
		data = append(data, make([]float64, cols+padding*2)...)
	}

	return mat.NewDense(rows+padding*2, cols+padding*2, data)
}

// output -> img_sections * filter_weights.T()
func Convolute(input *mat.Dense, filters *mat.Dense, stride, padding int, inputRows int, inputCols int) (*mat.Dense, int, int) {
	data := addPadding(input, padding)
	_, filterCols := filters.Dims()
	filterSize := int(math.Sqrt(float64(filterCols)))
	imageSections, sectionRows, sectionCols := extractImageSections(data, filterSize, stride, inputRows, inputCols)

	output := &mat.Dense{}
	output.Mul(imageSections, filters.T())
	return output, sectionRows, sectionCols
}

func ConvAndFcPredict(fc *Fc.NeuralLayer, conv *ConvolutionLayer, input []float64, inputRows, inputCols int) []float64 {
	convolution, _, _ := Convolute(mat.NewDense(1, len(input), input), conv.Filters, 1, 0, inputRows, inputCols)
	filterOutFlatten := mat.NewDense(1, len(convolution.RawMatrix().Data), convolution.RawMatrix().Data)
	return fc.Predict(filterOutFlatten).RawMatrix().Data
}

func ConvAndFcFit(alpha float64, fc *Fc.NeuralLayer, layer *ConvolutionLayer, input []float64, inputRows int,
	inputCols int, expected []float64, activeFunc Fc.ActiveFunc, derivFunc Fc.ActiveFunc) float64 {
	inputDense := mat.NewDense(1, len(input), input)
	expectedDense := mat.NewDense(1, len(expected), expected)

	// part of Convolute function
	data := addPadding(inputDense, 0)
	_, filterCols := layer.Filters.Dims()
	filterSize := int(math.Sqrt(float64(filterCols)))
	imageSections, _, _ := extractImageSections(data, filterSize, 1, inputRows, inputCols)

	filterOut := &mat.Dense{}
	filterOut.Mul(imageSections, layer.Filters.T())

	toReshapeRows, toReshapeCols := filterOut.Dims()
	// active func
	filterOutActiveFun := mat.NewDense(toReshapeRows, toReshapeCols, activeFunc(filterOut.RawMatrix().Data))

	filterOutFlatten := mat.NewDense(1, len(filterOutActiveFun.RawMatrix().Data), filterOutActiveFun.RawMatrix().Data)
	fcOut := fc.Predict(filterOutFlatten)

	fcDelta := &mat.Dense{}
	fcDelta.Sub(fcOut, expectedDense)

	filterDelta := &mat.Dense{}
	filterDelta.Mul(fcDelta, fc.Neurons)

	// reshaping
	filtersReshaped := mat.NewDense(toReshapeRows, toReshapeCols, filterDelta.RawMatrix().Data)
	// derive func
	filtersReshaped.MulElem(filtersReshaped, mat.NewDense(toReshapeRows, toReshapeCols, derivFunc(filterOut.RawMatrix().Data)))

	fcWeightDelta := &mat.Dense{}
	fcWeightDelta.Mul(fcDelta.T(), filterOutFlatten)
	fcWeightDelta.Scale(alpha, fcWeightDelta)
	fc.Neurons.Sub(fc.Neurons, fcWeightDelta)

	filterWeightDelta := &mat.Dense{}
	filterWeightDelta.Mul(filtersReshaped.T(), imageSections)
	filterWeightDelta.Scale(alpha, filterWeightDelta)
	layer.Filters.Sub(layer.Filters, filterWeightDelta)

	return 0.0
}

func ConvReluPoolFcPredict(fc *Fc.NeuralLayer, conv *ConvolutionLayer, input []float64, inputRows, inputCols int, maskSize int) []float64 {
	convolution, _, _ := Convolute(mat.NewDense(1, len(input), input), conv.Filters, 1, 0, inputRows, inputCols)

	// pooling
	pooledFlatten := mat.DenseCopyOf(convolution)

	return fc.Predict(pooledFlatten).RawMatrix().Data
}

func ConvReluPoolFcFit(alpha float64, fc *Fc.NeuralLayer, layer *ConvolutionLayer, input []float64, inputRows int,
	inputCols int, expected []float64, activeFunc Fc.ActiveFunc, derivFunc Fc.ActiveFunc, maskSize int) float64 {
	inputDense := mat.NewDense(1, len(input), input)
	expectedDense := mat.NewDense(1, len(expected), expected)

	// part of Convolute function
	data := addPadding(inputDense, 0)
	_, filterCols := layer.Filters.Dims()
	filterSize := int(math.Sqrt(float64(filterCols)))
	imageSections, _, _ := extractImageSections(data, filterSize, 1, inputRows, inputCols)

	filterOut := &mat.Dense{}
	filterOut.Mul(imageSections, layer.Filters.T())

	toReshapeRows, toReshapeCols := filterOut.Dims()
	// active func
	filterOutActiveFun := mat.NewDense(toReshapeRows, toReshapeCols, activeFunc(filterOut.RawMatrix().Data))
	filterOutFlatten := mat.NewDense(1, len(filterOutActiveFun.RawMatrix().Data), filterOutActiveFun.RawMatrix().Data)

	// ====== conv ends ======
	pooled, pooledMap := MaxPooling(filterOutActiveFun, maskSize)
	pooledFlatten := mat.NewDense(1, len(pooled.RawMatrix().Data), pooled.RawMatrix().Data)
	// ====== fc starts ======

	fcOut := fc.Predict(pooledFlatten)

	fcDelta := &mat.Dense{}
	fcDelta.Sub(fcOut, expectedDense)

	filterDelta := &mat.Dense{}
	filterDelta.Mul(fcDelta, fc.Neurons)

	// reshaping
	filtersReshaped := mat.NewDense(toReshapeRows, toReshapeCols, filterDelta.RawMatrix().Data)
	// derive func
	filtersReshaped.MulElem(filtersReshaped, mat.NewDense(toReshapeRows, toReshapeCols, derivFunc(filterOut.RawMatrix().Data)))

	// fc deltas
	fcWeightDelta := &mat.Dense{}
	fcWeightDelta.Mul(fcDelta.T(), filterOutFlatten)
	fcWeightDelta.Scale(alpha, fcWeightDelta)
	fc.Neurons.Sub(fc.Neurons, fcWeightDelta)

	// conv deltas
	filterWeightDelta := &mat.Dense{}
	filterWeightDelta.Mul(filtersReshaped.T(), imageSections)
	filterWeightDelta.Scale(alpha, filterWeightDelta)
	// inverse pooling (?)
	filterWeightDelta = ReshapePooling(filterWeightDelta, pooledMap, maskSize)
	layer.Filters.Sub(layer.Filters, filterWeightDelta)

	return 0.0
}

// returns pooled and binary representation of dense before pooling
func MaxPooling(data *mat.Dense, maskSize int) (*mat.Dense, *mat.Dense) {
	
}

func ReshapePooling(deltas, pooledMap *mat.Dense, maskSize int) *mat.Dense {

}