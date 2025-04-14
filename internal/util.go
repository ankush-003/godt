package internal

func GetUniqueValuesCount(data []Row, col int) (map[float32]int64, int64) {
	uniqueValues := make(map[float32]int64)
	var totalCount int64 = 0
	for _, row := range data {
		uniqueValues[row.Features[col]] += 1
		totalCount += 1
	}
	return uniqueValues, totalCount
}

// calculate Gini impurity for a given set of unique values
func CalculateGiniImpurity(
	data []Row,
	col int,
) float32 {
	uniqueMap, totalCount := GetUniqueValuesCount(data, col)
	if totalCount == 0 {
		return 0.0
	}

	floatUniqueCount := float32(totalCount)
	var giniScore float32 = 1.0
	for _, count := range uniqueMap {
		prob := float32(count) / floatUniqueCount
		giniScore -= prob * prob
	}
	return giniScore
}

func CalculateAverage(data []Row, col int) float32 {
	var sum float32 = 0.0
	for _, row := range data {
		sum += row.Features[col]
	}
	return sum / float32(len(data))
}

func CalculateMajority(data []Row, col int) float32 {
	uniqueMap, _ := GetUniqueValuesCount(data, col)
	var maxCount int64 = 0
	var majorityValue float32 = 0.0
	for value, count := range uniqueMap {
		if count > maxCount {
			maxCount = count
			majorityValue = value
		}
	}
	return majorityValue
}

func CalculateClassifierAccuracy(
	data Data,
	predictions []float32,
	target int,
) float32 {
	correctPredictions := data.RangeOverRow(func(i int, data *Row) *Row {
		if data.Features[target] == predictions[i] {
			return data
		}
		return nil
	})

	totalCount := correctPredictions.Size()
	return float32(totalCount) / float32(data.Size())
}
