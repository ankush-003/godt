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
	predictions := dt.Predict(*data)
	for i, prediction := range predictions {
		t.Logf("Actual: %v, Predicted: %v", data.Rows[i].Features[4], prediction)
		// assert.Equal(t, data.Rows[i].Features[4], prediction, "Predicted value should match the expected value")
	}
	accuracy := internal.CalculateClassifierAccuracy(*data, predictions, 4)
	t.Log("Accuracy:", accuracy)
	assert.GreaterOrEqual(t, accuracy, float32(0.95))
}
