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
