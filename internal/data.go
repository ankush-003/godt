package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Row struct {
	Features []float32
	Value    float32
	Pred     float32
}

type Data struct {
	// Headers is a slice of strings representing the column names
	Headers []string
	// Rows is a slice of Row structs representing the data
	Rows []Row
}

type dataOpts func(*Data) *Data

func FromCSV(path string) dataOpts {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	header, err := reader.Read()
	fmt.Println(header)
	if err != nil {
		panic(err)
	}
	rows := make([]Row, 0)
	for {
		row := make([]float32, len(header))
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		for i := 0; i < len(header); i++ {
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				continue
			}
			row[i] = float32(val)
		}
		rows = append(rows, Row{
			Features: row,
		})
	}

	return func(d *Data) *Data {
		size := len(d.Rows)
		d.Headers = header
		d.Rows = rows[:size]
		return d
	}
}

func NewData(size int, opts ...dataOpts) *Data {
	d := &Data{
		Rows: make([]Row, size),
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func (d *Data) DataWithRows(rows []Row) Data {
	return Data{
		Headers: d.Headers,
		Rows:    rows,
	}
}

func (d *Data) RangeOverCol(col int, f func(*Row) *Row) Data {
	res := Data{
		Headers: d.Headers,
		Rows:    make([]Row, 0),
	}
	for _, feature := range d.Rows {
		newRow := &Row{
			Features: feature.Features,
			Value:    feature.Value,
			Pred:     feature.Pred,
		}
		if updatedRow := f(newRow); updatedRow != nil {
			res.Rows = append(res.Rows, *updatedRow)
		}
	}
	return res
}

func (d *Data) Size() int {
	return len(d.Rows)
}

func (d *Data) Cols() int {
	if len(d.Rows) == 0 {
		return 0
	}
	// Assuming all rows have the same number of features
	return len(d.Rows[0].Features)
}

func (d *Data) GetColData(col int) []Row {
	coldData := make([]Row, d.Size())
	for i := 0; i < d.Size(); i++ {
		coldData[i].Features = []float32{d.Rows[i].Features[col]}
		coldData[i].Value = d.Rows[i].Value
		coldData[i].Pred = d.Rows[i].Pred
	}
	return coldData
}
