package internal

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func GenerateRegressionData(numRows int, numFeatures int) [][]float64 {
	data := make([][]float64, numRows)
	headers := make([]string, numFeatures)
	for i := 0; i < numFeatures; i++ {
		headers[i] = "feature_" + fmt.Sprint(i)
	}

	for i := 0; i < numRows; i++ {
		row := make([]float64, numFeatures)
		for j := 0; j < numFeatures-1; j++ {
			row[j] = rand.Float64() * 10                           // Random float between 0 and 10
			row[numFeatures-1] += row[j] * math.Pow(2, float64(j)) // Add polynomial term
		}
		data[i] = row
	}
	f, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write(headers)
	for _, row := range data {
		strRow := make([]string, numFeatures)
		for j, v := range row {
			strRow[j] = fmt.Sprintf("%f", v)
		}
		writer.Write(strRow)
	}
	return data
}

func GenerateNonLinear(numRows int) [][]float64 {
	data := make([][]float64, numRows)
	headers := []string{"feature_1", "value"}
	for i := 0; i < numRows; i++ {
		row := make([]float64, 2)
		row[0] = rand.Float64() * 10                                             // Random float between 0 and 10
		row[1] = 2*math.Sin(row[0]) + 2*math.Pow(row[0], 2) + rand.NormFloat64() // Non-linear function with noise
		data[i] = row
	}
	f, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write(headers)
	for _, row := range data {
		strRow := make([]string, len(row))
		for j, v := range row {
			strRow[j] = fmt.Sprintf("%f", v)
		}
		writer.Write(strRow)
	}
	return data
}
