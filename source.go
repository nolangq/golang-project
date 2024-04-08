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

/*
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
*/
func (phone *Cell) DeleteObject() {
	phone = &Cell{}
}

func main() {
	// Opens csv file
	file, err := os.Open("cells.csv")

	// Checks for error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes file
	defer file.Close()

	// Read from csv file
	reader := csv.NewReader(file)

	// Reads all values
	data, err := reader.ReadAll()

	// Checks for error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Prints all values
	for _, row := range data {
		fmt.Println(row)
	}

	// Testing struct methods
	myPhone := Cell{}
	myPhone.model = "iPhone"
	fmt.Println(myPhone.model)
	myPhone.DeleteObject()
	fmt.Println(myPhone.model)
}
