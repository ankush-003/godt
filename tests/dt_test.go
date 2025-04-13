package tests

import (
	"github.com/ankush-003/godt/decision_tree"
	"github.com/ankush-003/godt/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecisionTreeClassifier(t *testing.T) {
	data := internal.NewData(150, internal.FromCSV("iris.csv"))
	t.Log("Data Size:", data.Size())

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
		t.Logf("Actual: %v, Predicted: %v", test_rows[i].Features[4], prediction)
		assert.Equal(t, test_rows[i].Features[4], prediction, "Predicted value should match the expected value")
	}
}
