package MnistDBUtils

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type ParsedImages struct {
	MagicNumber, DataCount, Width, Height uint32
	Images [][]byte
}

func ParseImageFile(name string) ParsedImages {
	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir)
	data, err := ioutil.ReadFile(dir + "/" + name)
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

func GetInputVector(image []byte) []float64 {
	converted := make([]float64, len(image))
	for i := range image {
		converted[i] = float64(image[i]) / 255
	}
	return converted
}

func GetInputMatrix(images [][]byte) [][]float64 {
	converted := make([][]float64, len(images))
	for each := range converted {
		converted[each] = GetInputVector(images[each])
	}
	return converted
}