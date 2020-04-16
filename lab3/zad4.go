package lab3

import (
	fun "../FunctionsAndDerivatives"
	SI "../SIFullyConnected"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"
)

type ParsedLabels struct {
	MagicNumber, DataCount uint32
	Labels []byte
}

type ParsedImages struct {
	MagicNumber, DataCount, Width, Height uint32
	Images [][]byte
}

func Zad4() {
	// don't know how to get this value, thus its hard coded for now
	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir)
	rand.Seed(time.Now().UnixNano())
	// network creation
	alpha := 0.01
	network := SI.CreateNetwork(
		alpha,
		[]int{784, 40, 10},
		[]SI.ActiveFunc{fun.ReLuFunc, fun.ReLuFunc},
		[]SI.ActiveFunc{fun.ReLuFuncDeriv, fun.ReLuFuncDeriv},
		)

	trainLabels := parseLabelFile(dir + "/train-labels.idx1-ubyte")
	trainImages := parseImageFile(dir + "/train-images.idx3-ubyte")

	testLabels := parseLabelFile(dir + "/t10k-labels.idx1-ubyte")
	testImages := parseImageFile(dir + "/t10k-images.idx3-ubyte")

	netErr := 0.0
	for i, limit := 0, 100; i < limit; i++{
		for set := range trainLabels.Labels {
			netErr = network.Fit(
				getExpectedVector(trainLabels.Labels[set]),
				getInputVector(trainImages.Images[set]),
				)
		}

		var prediction []float64
		errorCounter := 0
		for set := range testLabels.Labels {
			prediction = network.Predict(getInputVector(testImages.Images[set]))
			if testLabels.Labels[set] != getOutputLabel(prediction) {
				errorCounter++
			}
		}

		if i%10 == 0 {
				fmt.Printf(
				"%d iteration, network error: %f network score: %d/%d\n",
				i,
				netErr,
				int(testLabels.DataCount) - errorCounter,
				testLabels.DataCount,
				)
		}

	}
}

func parseLabelFile(path string) ParsedLabels {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	labels := ParsedLabels{
		MagicNumber: 0,
		DataCount:   0,
		Labels:      nil,
	}
	var offset, BTR, copiedBytes uint32
	offset = 0
	BTR = 4 // Bytes To Read

	labels.MagicNumber = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR
	labels.DataCount = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR

	labels.Labels = make([]byte, labels.DataCount)
	copiedBytes = uint32(copy(labels.Labels, data[offset : offset + labels.DataCount]))
	if copiedBytes != labels.DataCount {
		log.Panic("Did not read enough data")
	}
	return labels
}

func parseImageFile(path string) ParsedImages {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	images := ParsedImages{
		MagicNumber: 0,
		DataCount:   0,
		Width:       0,
		Height:      0,
		Images:      nil,
	}

	var offset, BTR uint32
	offset = 0
	BTR = 4 // Bytes To Read

	images.MagicNumber = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR
	images.DataCount = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR
	images.Width = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR
	images.Height = binary.LittleEndian.Uint32(swapBytes(data[offset : offset + BTR]))
	offset += BTR

	images.Images = make([][]byte, images.DataCount)
	for image := range images.Images {
		imageLen := images.Height * images.Width
		images.Images[image] = make([]byte, imageLen)
		copy(images.Images[image], data[offset : offset + imageLen])
		offset += imageLen
	}

	return images
}

func swapBytes(byt []byte) []byte {
	bytR := make([]byte, len(byt))
	for i := range byt {
		bytR[i] = byt[len(byt) - i - 1]
	}
	return bytR
}

func getExpectedVector(label byte) []float64 {
	expected := make([]float64, 10)
	expected[label] = 1.0
	return expected
}

func getInputVector(image []byte) []float64 {
	converted := make([]float64, len(image))
	for i := range image {
		converted[i] = float64(image[i]) / 255
	}
	return converted
}

func getOutputLabel(prediction []float64) byte {
	var label byte
	label = 0
	for i := range prediction {
		if prediction[label] < prediction[i] {
			label = byte(i)
		}
	}
	return label
}