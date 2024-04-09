package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	oem               string
	model             string
	launchAnnounced   int
	launchStatus      string
	bodyDimensions    string
	bodyWeight        float64
	bodySim           string
	displayType       string
	displaySize       float64
	displayResolution string
	featuresSensors   string
	platformOs        string
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
	records, err := reader.ReadAll()

	// Checks for error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Create data structure for storage (map)
	cellMap := make(map[int]Cell)
	index := 0

	// Process each row
	for _, record := range records {
		// Replace missing or "-" values with empty string
		for i := range record {
			if record[i] == "-" {
				record[i] = "" // Replace "-" with empty string
			}
		}

		// Convert launch announced to integer
		launchYear := extractYear(record[2])

		// Convert body weight to float
		bodyWeight := parseFloat(record[5])

		// Convert display size to float
		displaySize := parseFloat(record[8])

		// Create a new object
		cell := Cell{
			oem:               record[0],
			model:             record[1],
			launchAnnounced:   launchYear,
			launchStatus:      record[3],
			bodyDimensions:    record[4],
			bodyWeight:        bodyWeight,
			bodySim:           record[6],
			displayType:       record[7],
			displaySize:       displaySize,
			displayResolution: record[9],
			featuresSensors:   record[10],
			platformOs:        record[11],
		}

		// Store the object in the map with the current index
		cellMap[index] = cell

		// Increment index
		index++
	}

	// Testing
	indexToLookup := 2

	if cell, ok := cellMap[indexToLookup]; ok {
		fmt.Printf("Cell details for index %d:\n", indexToLookup)
		fmt.Printf("OEM: %s\n", cell.oem)
		fmt.Printf("Launch Announced: %d\n", cell.launchAnnounced)
		fmt.Printf("Body Weight: %.2f\n", cell.bodyWeight)
		fmt.Printf("Display Size: %.2f\n", cell.displaySize)
	} else {
		fmt.Printf("Cell with index %d not found\n", indexToLookup)
	}
}

func extractYear(dateStr string) int {
	date, err := time.Parse("January 2, 2006", dateStr)
	if err != nil {
		return 0
	}
	return date.Year()
}

// Change string to float
func parseFloat(sizeStr string) float64 {
	parts := strings.Fields(sizeStr)
	if len(parts) > 0 {
		size, err := strconv.ParseFloat(parts[0], 64)
		if err == nil {
			return size
		}
	}
	return 0.0
}

// Setter and getter functions for Cell
func (phone Cell) getOem() string {
	return phone.oem
}
func (phone Cell) setOem(tempOem string) {
	phone.oem = tempOem
}
func (phone Cell) getModel() string {
	return phone.model
}
func (phone Cell) setModel(tempModel string) {
	phone.model = tempModel
}
func (phone Cell) getLaunchAnnounced() int {
	return phone.launchAnnounced
}
func (phone Cell) setLaunchAnnounced(tempLaunchAnnounced int) {
	phone.launchAnnounced = tempLaunchAnnounced
}
func (phone Cell) getLaunchStatus() string {
	return phone.launchStatus
}
func (phone Cell) setLaunchStatus(tempLaunchStatus string) {
	phone.launchStatus = tempLaunchStatus
}
func (phone Cell) getBodyDimensions() string {
	return phone.bodyDimensions
}
func (phone Cell) setBodyDimensions(tempbodyDimensions string) {
	phone.bodyDimensions = tempbodyDimensions
}
func (phone Cell) getBodyWeight() float64 {
	return phone.bodyWeight
}
func (phone Cell) setBodyWeight(tempBodyWeight float64) {
	phone.bodyWeight = tempBodyWeight
}
func (phone Cell) getBodySim() string {
	return phone.bodySim
}
func (phone Cell) setBodySim(tempBodySim string) {
	phone.bodySim = tempBodySim
}
func (phone Cell) getDisplayType() string {
	return phone.displayType
}
func (phone Cell) setDisplayType(tempDisplayType string) {
	phone.displayType = tempDisplayType
}
func (phone Cell) getDisplaySize() float64 {
	return phone.displaySize
}
func (phone Cell) setDisplaySize(tempDisplaySize float64) {
	phone.displaySize = tempDisplaySize
}
func (phone Cell) getDisplayResolution() string {
	return phone.displayResolution
}
func (phone Cell) setDisplay_Resolution(tempDisplayResolution string) {
	phone.displayResolution = tempDisplayResolution
}
func (phone Cell) getFeaturesSensors() string {
	return phone.featuresSensors
}
func (phone Cell) setFeaturesSensors(tempFeaturesSensors string) {
	phone.featuresSensors = tempFeaturesSensors
}
func (phone Cell) gePlatformOs() string {
	return phone.platformOs
}
func (phone Cell) setPlatformOs(tempPlatformOs string) {
	phone.platformOs = tempPlatformOs
}
