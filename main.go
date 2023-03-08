package main

import (
	"encoding/csv"
	"os"
)

type CSV struct {
	header []string
	data   [][]string
}

func main() {
	var csv CSV
	csv.header = []string{"No", "Name", "Phone Number"}
	csv.data = [][]string{
		{"1", "John", "123-456-7890"},
		{"2", "Jane", "456-789-0123"},
		{"3", "Bob", "789-012-3456"},
	}

	err := createCSVFile("data.csv", header, data)
	if err != nil {
		panic(err)
	}
}

// createCSVFile creates a new CSV file with the specified filename, header row, and data rows
func createCSVFile(filename string, header []string, data [][]string) error {
	// Create the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the header row to the CSV file
	err = writeCSVRow(writer, header)
	if err != nil {
		return err
	}

	// Write the data rows to the CSV file
	for _, row := range data {
		err = writeCSVRow(writer, row)
		if err != nil {
			return err
		}
	}

	// Flush the CSV writer
	writer.Flush()

	return nil
}

// writeCSVRow writes a single row to the specified CSV writer
func writeCSVRow(writer *csv.Writer, row []string) error {
	err := writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}
