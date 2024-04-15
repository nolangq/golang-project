package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

		launchYearString := extractFirstFourDigits(record[2])
		launchYear := stringToInt(launchYearString)

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
	// Call the countPhonesWithOneSensor function
	count := countPhonesWithOneSensor(cellMap)

	// Print the count of phones with only one feature sensor
	fmt.Printf("Number of phones with only one feature sensor: %d\n", count)

	results := findAnnouncedAndReleasedInDifferentYears(cellMap)
	fmt.Println("Phones announced in one year and released in another year:")
	for _, phone := range results {
		fmt.Printf("OEM: %s, Model: %s\n", phone.oem, phone.model)
	}

	highestAvgOEM := findHighestAvgBodyWeightOEM(cellMap)
	fmt.Printf("OEM with the highest average body weight: %s\n", highestAvgOEM)

	indexToLookup := 3

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

// Report Functions
// Calculates highest average body weight
func countPhonesWithOneSensor(cells map[int]Cell) int {
	count := 0

	for _, cell := range cells {
		// Split the featuresSensors string by comma to count the number of sensors
		sensors := strings.Split(cell.featuresSensors, ",")

		// Remove leading and trailing spaces from each sensor name
		for i := range sensors {
			sensors[i] = strings.TrimSpace(sensors[i])
		}

		// Check if exactly one sensor is present
		if len(sensors) == 1 && sensors[0] != "" {
			count++
		}
	}

	return count
}

func findHighestAvgBodyWeightOEM(cellMap map[int]Cell) string {
	// Create map to store cumulative body weights and count of phones per OEM
	weightSum := make(map[string]float64)
	count := make(map[string]int)

	for _, cell := range cellMap {
		weightSum[cell.oem] += cell.bodyWeight
		count[cell.oem]++
	}

	var highestAvgOEM string
	var maxAvgWeight float64

	// Calculate average body weight for each OEM
	for oem, sum := range weightSum {
		avgWeight := sum / float64(count[oem])
		if avgWeight > maxAvgWeight {
			maxAvgWeight = avgWeight
			highestAvgOEM = oem
		}
	}

	return highestAvgOEM
}

func findAnnouncedAndReleasedInDifferentYears(cellMap map[int]Cell) []Cell {
	var result []Cell

	for _, cell := range cellMap {
		if cell.launchStatus != "" {
			releaseYear := extractReleaseYear(cell.launchStatus)
			if releaseYear != 0 && releaseYear != cell.launchAnnounced {
				result = append(result, cell)
			}
		}
	}

	return result
}

func extractReleaseYear(launchStatus string) int {
	// Split the launch status by spaces and commas
	parts := strings.FieldsFunc(launchStatus, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, part := range parts {
		if len(part) == 4 {
			if year, err := strconv.Atoi(part); err == nil {
				return year
			}
		}
	}

	return 0 // Return 0 if release year not found or invalid format
}

func extractFirstFourDigits(input string) string {
	var digits []rune
	digitCount := 0

	// Iterate over each character
	for _, char := range input {
		// Check if character is a digit
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			digitCount++
			if digitCount == 4 {
				break
			}
		} else {
			// Reset digitCount
			digits = nil
			digitCount = 0
		}
	}

	// Convert the collected digits to a string
	return string(digits)
}

func stringToInt(s string) int {
	// Convert the string to an integer
	num, _ := strconv.Atoi(s)
	return num
}

func floatToString(value float64) string {
	// Convert float64 to string
	return strconv.FormatFloat(value, 'f', -1, 64)
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

// Seven class methods
// Display phone's details
func (c Cell) DisplayDetails() {
	fmt.Printf("OEM: %s\n", c.oem)
	fmt.Printf("Model: %s\n", c.model)
	fmt.Printf("Launch Announced: %d\n", c.launchAnnounced)
	fmt.Printf("Launch Status: %s\n", c.launchStatus)
	fmt.Printf("Body Dimensions: %s\n", c.bodyDimensions)
	fmt.Printf("Body Weight: %.2f\n", c.bodyWeight)
	fmt.Printf("Body SIM: %s\n", c.bodySim)
	fmt.Printf("Display Type: %s\n", c.displayType)
	fmt.Printf("Display Size: %.2f\n", c.displaySize)
	fmt.Printf("Display Resolution: %s\n", c.displayResolution)
	fmt.Printf("Features & Sensors: %s\n", c.featuresSensors)
	fmt.Printf("Platform OS: %s\n", c.platformOs)
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
