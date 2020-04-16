package MnistDBUtils

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type ParsedLabels struct {
	MagicNumber, DataCount uint32
	Labels []byte
}

func ParseLabelFile(name string) ParsedLabels {
	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir)
	data, err := ioutil.ReadFile(dir + "/" + name)
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

func GetExpectedVector(label byte) []float64 {
	expected := make([]float64, 10)
	expected[label] = 1.0
	return expected
}

func GetOutputLabel(prediction []float64) byte {
	var label byte
	label = 0
	for i := range prediction {
		if prediction[label] < prediction[i] {
			label = byte(i)
		}
	}
	return label
}