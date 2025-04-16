package internal

import "sort"

type Argsort struct {
	Indices []int
	data    []Row
	col     int
}

func (a Argsort) Len() int {
	return len(a.data)
}

func (a Argsort) Swap(i, j int) {
	a.Indices[i], a.Indices[j] = a.Indices[j], a.Indices[i]
	a.data[i], a.data[j] = a.data[j], a.data[i]
}

func (a Argsort) Less(i, j int) bool {
	return a.data[i].Features[a.col] < a.data[j].Features[a.col]
}

func ArgSortRows(data []Row, col int) []int {
	indices := make([]int, len(data))
	for i := range indices {
		indices[i] = i
	}
	sort.Sort(Argsort{Indices: indices, data: data, col: col})
	return indices
}
