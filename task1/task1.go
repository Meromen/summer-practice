package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)
// 9200 9299

const (
	IMAGE_WIDTH      = 170
	IMAGE_HEIGHT     = 100
	A_VALUE_MLP_CF   = 0.6  // 0 - 1
	MAX_VALUE_DIV_CF = 20.0 // > 1
)

func main() {
	file, err := ioutil.ReadFile(`task1Data.txt`)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	maxValue := 0.0
	aValue := 0.0

	for _, line := range lines {
		lineValues := strings.Split(line, " ")

		for _, lineValue := range lineValues {
			numericValue, err := strconv.ParseFloat(lineValue, 64)
			if err != nil {
				panic(err)
			}
			if maxValue < numericValue {
				maxValue = numericValue
			}
		}
	}

	aValue = A_VALUE_MLP_CF * maxValue

	lineImg := image.NewRGBA(image.Rect(0, 0 , IMAGE_WIDTH, IMAGE_HEIGHT))
	logImg :=  image.NewRGBA(image.Rect(0, 0 , IMAGE_WIDTH, IMAGE_HEIGHT))


	for lineIndex, line := range lines {
		lineValues := strings.Split(line, " ")

		for valueIndex, lineValue := range lineValues {
			numericValue, err := strconv.ParseFloat(lineValue, 64)
			if err != nil {
				panic(err)
			}

			lineColorValue := (numericValue / (maxValue / MAX_VALUE_DIV_CF)) * 255
			logColorValue := 255 * (math.Log10(((aValue - 1) / maxValue) * numericValue + 1) / math.Log10(aValue))

			lineImg.Set(valueIndex, lineIndex, color.RGBA{
				R: uint8(lineColorValue),
				G: uint8(lineColorValue),
				B: uint8(lineColorValue),
				A: 255,
			})

			logImg.Set(valueIndex, lineIndex, color.RGBA{
				R: uint8(logColorValue),
				G: uint8(logColorValue),
				B: uint8(logColorValue),
				A: 255,
			})
		}
	}

	outputFile, err := os.OpenFile("task1Line.jpg", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, lineImg, nil)
	if err != nil {
		panic(err)
	}

	outputFile, err = os.OpenFile("task1Log.jpg", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, logImg, nil)
	if err != nil {
		panic(err)
	}
}