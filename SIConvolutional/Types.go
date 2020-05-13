package SIConvolutional

import "gonum.org/v1/gonum/mat"

type LayerType int

const (
	CONV   = 0 // convolutional
	ACTIVE = 1 // layer with activation function
	POOL   = 2 // pooling layer
	FC     = 3 // fully connected layer
)

type Layer struct {
	Filters []*mat.Dense
	Type  LayerType
}

type Convolution struct {
	Mask    interface{}
	Filters interface{}
	Stride  int
	Padding int
}

type ConvolutionalNet struct {
	Layers Layer
}
