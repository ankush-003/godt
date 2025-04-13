package decision_tree

import (
	"github.com/ankush-003/godt/internal"
	"github.com/rs/zerolog/log"
)

type BasicNode struct {
	Children  []*DTNode
	Depth     int
	Variance  float32
	Criterion float32
	Feature   int
}

type DTNode struct {
	Node         *BasicNode
	impurity     float32
	Data         internal.Data
	TargetColumn int
	MaxDepth     int
}

func NewDTNode(
	depth int,
	maxDepth int,
	targetColumn int,
) *DTNode {
	return &DTNode{
		Node: &BasicNode{
			Children: nil,
			Depth:    depth,
		},
		Data:         internal.Data{},
		TargetColumn: targetColumn,
		MaxDepth:     maxDepth,
	}
}

func (d *DTNode) DataLen() float32 {
	return float32(len(d.Data.Rows))
}

func (d *DTNode) IsLeaf() bool {
	return len(d.Node.Children) == 0
}

func (d *DTNode) CalculateGain(info *FeatureSplitInfo) float32 {
	gain := d.impurity
	length := d.DataLen()
	if length == 0 {
		return 0.0
	}

	for i, feature := range info.splits {
		gain -= info.impurities[i] * (float32(len(feature)) / length)
	}
	return gain
}

type FeatureSplitInfo struct {
	splits     [][]internal.Row
	impurities []float32
	Gain       float32
	Feature    int
	Criterion  float32
}

func (d *DTNode) FindBestSplitForColumn(column int) *FeatureSplitInfo {
	uniqueMap, _ := internal.GetUniqueValuesCount(d.Data.Rows, column)
	uniqueCount := float32(len(uniqueMap))
	if uniqueCount == 1 {
		log.Info().
			Int("depth", d.Node.Depth).
			Int("Column", column).
			Msg("unique count is 1")
		return nil
	}

	log.Info().
		Int("depth", d.Node.Depth).
		Int("Column", column).
		Float32("unique_count", uniqueCount).
		Msg("unique count")
	data := d.Data
	bestSplit := &FeatureSplitInfo{
		Gain: 0.0,
	}
	for feature, _ := range uniqueMap {
		left := data.RangeOverCol(column, func(row *internal.Row) *internal.Row {
			if row.Features[column] <= feature {
				return row
			}
			return nil
		})
		right := data.RangeOverCol(column, func(row *internal.Row) *internal.Row {
			if row.Features[column] > feature {
				return row
			}
			return nil
		})
		left_impurity := internal.CalculateGiniImpurity(left.Rows, d.TargetColumn)
		right_impurity := internal.CalculateGiniImpurity(right.Rows, d.TargetColumn)

		log.Info().
			Int("depth", d.Node.Depth).
			Int("Column", column).
			Float32("left_impurity", left_impurity).
			Float32("right_impurity", right_impurity).
			Float32("feature", feature).
			Msg("impurity")

		newSplitFeature := &FeatureSplitInfo{
			splits:     [][]internal.Row{left.Rows, right.Rows},
			impurities: []float32{left_impurity, right_impurity},
			Gain:       0.0,
			Feature:    column,
			Criterion:  feature,
		}
		gain := d.CalculateGain(newSplitFeature)
		if gain > bestSplit.Gain {
			bestSplit = newSplitFeature
			log.Info().
				Int("depth", d.Node.Depth).
				Int("Column", column).
				Float32("gain", gain).
				Float32("criterion", feature).
				Msg("best gain")
		}
	}

	return bestSplit
}

func (d *DTNode) FindBestSplit() *FeatureSplitInfo {
	var bestSplit *FeatureSplitInfo = nil
	for i := 0; i < d.Data.Cols(); i++ {
		split := d.FindBestSplitForColumn(i)
		if split == nil {
			continue
		}
		if bestSplit == nil {
			bestSplit = split
			continue
		}
		if split.Gain > bestSplit.Gain {
			bestSplit = split
		}
	}
	return bestSplit
}

func (d *DTNode) Fit() {
	if d.Node.Depth >= d.MaxDepth {
		return
	}
	log.Info().
		Int("depth", d.Node.Depth).
		Float32("impurity", d.impurity).
		Int("data_size", d.Data.Size()).
		Msg("Fitting DTNode")

	SplitInfo := d.FindBestSplit()
	if SplitInfo == nil {
		return
	}
	d.Node.Criterion = SplitInfo.Criterion
	d.Node.Feature = SplitInfo.Feature

	log.Info().
		Int("depth", d.Node.Depth).
		Msg("Creating branches")
	d.CreateBranches(SplitInfo)
}

func (d *DTNode) CreateBranches(
	split *FeatureSplitInfo,
) {
	if len(split.splits) == 0 {
		return
	}
	leftChild := NewDTNode(d.Node.Depth+1, d.MaxDepth, d.TargetColumn)
	rightChild := NewDTNode(d.Node.Depth+1, d.MaxDepth, d.TargetColumn)
	leftChild.Data = d.Data.DataWithRows(split.splits[0])
	rightChild.Data = d.Data.DataWithRows(split.splits[1])
	leftChild.impurity = split.impurities[0]
	rightChild.impurity = split.impurities[1]
	d.Node.Children = append(d.Node.Children, leftChild)
	d.Node.Children = append(d.Node.Children, rightChild)
	leftChild.Fit()
	rightChild.Fit()
}

func (d *DTNode) NodeAverage() float32 {
	return internal.CalculateAverage(d.Data.Rows, d.TargetColumn)
}

func (d *DTNode) NodeMajority() float32 {
	return internal.CalculateMajority(d.Data.Rows, d.TargetColumn)
}

func (d *DTNode) Traverse(
	row internal.Row,
) float32 {
	if d.IsLeaf() {
		return d.NodeMajority()
	}
	if row.Features[d.Node.Feature] <= d.Node.Criterion {
		return d.Node.Children[0].Traverse(row)
	}
	return d.Node.Children[1].Traverse(row)
}
