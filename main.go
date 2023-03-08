package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// Define the header row of the CSV file
	header := []string{"No", "Name", "Phone Number"}

	// Define the data rows of the CSV file
	data := [][]string{
		{"1", "John", "123-456-7890"},
		{"2", "Jane", "456-789-0123"},
		{"3", "Bob", "789-012-3456"},
	}

	// Create the CSV file
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the header row to the CSV file
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	// Write the data rows to the CSV file
	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
	}

	// Flush the CSV writer
	writer.Flush()
}
