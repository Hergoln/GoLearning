package lab3

import (
	"../LegacySIFullyConnected"
	"fmt"
	"io"
	"io/ioutil"
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

/*
	Q: Która sieć potrzebuje więcej czasu, żeby nauczyć się rozpoznawania kolorów?
	A: Sieć głęboka potrzebuje więcej czasu, może to wynikać z faktu, że głęboka uczy się
	wzorców a nie konkretnych przypadków.
*/

func Zad3() {
	// 3.3
	LIMIT := 500
	STEP := 10
	network := LegacySIFullyConnected.ConstructRandomNetwork(3, 4)
	network.AppendRandomLayerWithActiveFunc(4, nil, nil)

	resp, err := http.Get("http://pduch.iis.p.lodz.pl/PSI/training_colors.txt")
	if err != nil {
		fmt.Println(err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	trainingLines := parseTestingFiles(resp.Body)

	resp, err = http.Get("http://pduch.iis.p.lodz.pl/PSI/test_colors.txt")
	if err != nil {
		fmt.Println(err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	testingLines := parseTestingFiles(resp.Body)
	alpha := .01

	errorCounter := 0
	errNet := 0.
	for i := 0; i < LIMIT ; i++ {
		errorCounter = 0
		errNet = 0.
		for line := range trainingLines {
			var expected [4]float64
			expected[trainingLines[line].expected - 1] = 1
			errNet += network.StudyActiveFunc(alpha, expected[:], trainingLines[line].input[:], ReLu)
		}

		var prediction []float64

		for line := range testingLines {
			var expectedOutput [4]float64
			expectedOutput[testingLines[line].expected - 1] = 1.
			prediction = network.PredictActiveFunc(testingLines[line].input[:], ReLu)
			indPre, _ := getColorAndIndex(prediction)
			indExp, _ := getColorAndIndex(expectedOutput[:])
			if indPre != indExp {
				errorCounter++
			}
		}

		if i%STEP == 0 {
			fmt.Printf("\n============================ %d ============================\n", i)
			fmt.Printf("(DEEP)Score: %d/%d; Error: %f\n", errorCounter, len(testingLines), errNet)
			network.DisplayNet()
		}

		//if errorCounter < int(float64(len(testingLines)) * 0.1) {
		if errorCounter == 0 {
			fmt.Printf("\n(DEEP)Good score(%d/%d), %d iterations\n", errorCounter, len(testingLines), i)
			return
		}
	}

	fmt.Printf("\n============================ %d ============================\n", LIMIT)
	fmt.Printf("(DEEP)Score: %d/%d; Error: %f\n", errorCounter, len(testingLines), errNet)
	network.DisplayNet()
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