package decision_tree

import (
	"fmt"

	"github.com/ankush-003/godt/internal"
	"github.com/rs/zerolog/log"
)

type DecisionTree interface {
	// Fit the model to the training data
	Fit()
	// Predict the target value for a given input
	Predict(input internal.Data) float32
}

type DecisionTreeClassifier struct {
	Root            *DTNode
	MaxDepth        int
	MinSamplesSplit int
	TargetColumn    int
	Data            internal.Data
}

func NewDecisionTreeClassifier(
	data internal.Data,
	maxDepth int,
	minSamplesSplit int,
	targetColumn int,
) *DecisionTreeClassifier {
	log.Info().
		Int("max_depth", maxDepth).
		Int("min_samples_split", minSamplesSplit).
		Int("target_column", targetColumn).
		Int("data_size", data.Size()).
		Int("Columns", data.Cols()).
		Strs("headers", data.Headers).
		Msg("Creating Decision Tree Regressor")

	return &DecisionTreeClassifier{
		Root:            nil,
		MaxDepth:        maxDepth,
		MinSamplesSplit: minSamplesSplit,
		TargetColumn:    targetColumn,
		Data:            data,
	}
}

func (d *DecisionTreeClassifier) Fit() {
	log.Info().Msg("Fitting Decision Tree Classifier")
	d.Root = NewDTNode(0, d.MaxDepth, d.TargetColumn)
	d.Root.Data = d.Data
	fmt.Println("Data Size:", d.Data.Size())
	fmt.Println("Data Columns:", d.Data.Cols())
	fmt.Println("Data Headers:", d.Data.Headers)
	fmt.Println("Data Rows:", d.Data.Rows[0].Features)

	d.Root.impurity = internal.CalculateGiniImpurity(d.Data.Rows, d.TargetColumn)
	d.Root.Fit()
}

func (d *DecisionTreeClassifier) Predict(input internal.Data) []float32 {
	log.Info().Msg("Predicting using Decision Tree Classifier")
	if d.Root == nil {
		log.Error().Msg("Decision Tree Classifier not fitted yet")
		return nil
	}

	predictions := make([]float32, len(input.Rows))
	for i, row := range input.Rows {
		predictions[i] = d.Root.Traverse(row)
	}
	return predictions
}
