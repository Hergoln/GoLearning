package lab2

import (
	"../LegacySIFullyConnected"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	Red		= 0
	Green	= 1
	Blue	= 2
	Yellow 	= 3
)
var ColorNames = map[int]string{
	Red: "Red",
	Green: "Green",
	Blue: "Blue",
	Yellow: "Yellow",
}

type TestFileLine struct {
	input [3]float64
	expected int
}

func Zad4() {
	// 2.4
	network := LegacySIFullyConnected.ConstructRandomNetwork(3, 4).Layers[0]
	resp, err := http.Get("http://pduch.iis.p.lodz.pl/PSI/training_colors.txt")
	if err != nil {
		log.Println(err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	trainingLines := parseTestingFiles(resp.Body)

	resp, err = http.Get("http://pduch.iis.p.lodz.pl/PSI/test_colors.txt")
	if err != nil {
		log.Println(err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	testingLines := parseTestingFiles(resp.Body)
	alpha := 0.01

	for learningIterations := 0; ; learningIterations++ {
		errorCounter := 0

		for line := range trainingLines {
			var expectedOutput [4]float64
			expectedOutput[trainingLines[line].expected - 1] = 1
			network.Study(alpha, expectedOutput[:], trainingLines[line].input[:])
		}

		var prediction []float64
		for line := range testingLines {
			var expectedOutput [4]float64
			expectedOutput[testingLines[line].expected - 1] = 1.
			prediction = network.Predict(testingLines[line].input[:])
			indPre, _ := getColorAndIndex(prediction)
			indExp, _ := getColorAndIndex(expectedOutput[:])
			if indPre != indExp {
				errorCounter++
			}
		}
		if errorCounter == 0 {
			fmt.Printf("\n2.4\n(NORM)Perfect score, %d iterations\n", learningIterations)
			return
		}
	}
}

func parseTestingFiles(reader io.Reader) []TestFileLine  {
	bytesSlice, _ := ioutil.ReadAll(reader)
	text := string(bytesSlice)

	lines := strings.Split(text, "\n")
	toReturn := make([]TestFileLine, len(lines))
	var parsed TestFileLine
	for l := range lines {
		parts := strings.Split(lines[l], " ")
		parsed.input[0], _ = strconv.ParseFloat(parts[0], 64)
		parsed.input[1], _ = strconv.ParseFloat(parts[1], 64)
		parsed.input[2], _ = strconv.ParseFloat(parts[2], 64)
		temp, _ := strconv.ParseInt(strings.TrimSpace(parts[3]), 10, 32)
		parsed.expected = int(temp)
		toReturn[l] = parsed
	}
	return toReturn
}

func getColorAndIndex(prediction []float64) (int, string) {
	ind := 0
	for i := range prediction {
		if prediction[ind] <= prediction[i] {
			ind = i
		}
	}
	return ind, ColorNames[ind]
}