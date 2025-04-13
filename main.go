package main

import (
	"github.com/ankush-003/godt/decision_tree"
	"github.com/ankush-003/godt/internal"
	"github.com/rs/zerolog/log"
)

func main() {
	data := internal.NewData(150, internal.FromCSV("./tests/iris.csv"))

	dt := decision_tree.NewDecisionTreeClassifier(
		*data,
		12,
		10,
		4,
	)
	dt.Fit()
	test_rows := []internal.Row{
		{
			Features: []float32{5.4, 3.9, 1.7, 0.4, 0},
		},
		{
			Features: []float32{4.6, 3.4, 1.4, 0.3, 0},
		},
		{
			Features: []float32{5.9, 3.0, 5.1, 1.8, 2},
		},
	}
	test_Data := data.DataWithRows(test_rows)
	predictions := dt.Predict(test_Data)
	for i, prediction := range predictions {
		log.Info().
			Float32("actual", test_rows[i].Features[4]).
			Float32("predicted", prediction).
			Msg("Prediction")
	}
}
