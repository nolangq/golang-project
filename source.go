package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Cell struct {
	oem                string
	model              string
	launch_announced   string
	launch_status      string
	body_dimensions    string
	body_weight        string
	body_sim           string
	display_type       string
	display_size       string
	display_resolution string
	features_sensors   string
	platform_os        string
}

func (phone Cell) ToString() string {

}
func (phone Cell) CalculateMean() string {

}
func (phone Cell) CalculateMedian() string {

}
func (phone Cell) CalculateStandardDeviation() string {

}
func (phone Cell) CalculateMode() string {

}
func (phone Cell) ListUniqueValues() string {

}
func (phone Cell) AddDataAndInput() string {

}
func (phone Cell) DeleteObject() {

}

func main() {
	file, err := os.Open("cells.csv")

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}
}
