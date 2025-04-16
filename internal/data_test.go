package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFromCSV(t *testing.T) {
	// Test the FromCSV function with a sample CSV file
	// Create a temporary CSV file for testing
	csvData := `feature1,feature2,value
1.0,2.0,3.0
4.0,5.0,6.0`

	tmpFile, err := os.CreateTemp("", "testdata.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temp file after the test
	_, err = tmpFile.WriteString(csvData)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	// Call the FromCSV function
	d := NewData(2, FromCSV(tmpFile.Name()))
	// Check if the headers are set correctly
	expectedHeaders := []string{"feature1", "feature2", "value"}
	assert.Equal(t, expectedHeaders, d.Headers, "Headers should match the CSV header")
	assert.Equal(t, 2, len(d.Rows), "Data length should match the number of rows in CSV")
	assert.Equal(t, float32(1), d.Rows[0].Features[0], "First feature should be 1")
}

func TestArgSort(t *testing.T) {
	GenerateNonLinear(10)
	defer os.Remove("data.csv")
	d := NewData(10, FromCSV("data.csv"))

	t.Log("Before sorting:")
	for i := 0; i < len(d.Rows); i++ {
		t.Log(d.Rows[i].Features[0])
	}

	row_copy := make([]Row, len(d.Rows))
	copy(row_copy, d.Rows)
	sortedIndices := ArgSortRows(row_copy, 0)

	t.Log("After sorting:")
	t.Log("Sorted Indices:", sortedIndices)
	for i := 0; i < len(sortedIndices)-1; i++ {
		t.Log(d.Rows[sortedIndices[i]].Features[0])
		assert.LessOrEqual(t, d.Rows[sortedIndices[i]].Features[0], d.Rows[sortedIndices[i+1]].Features[0], "Rows should be sorted in ascending order")
	}
}
