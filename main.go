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

	// Test 1
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

	// Cell functions in use
	// Creating temp phone
	var myPhone = Cell{}
	// Plugging random values
	myPhone.setOem("Apple")
	myPhone.setModel("iPhone 12")
	myPhone.setBodyDimensions("145 x 56 x 23 mm (5.71 x 2.20 x 0.91 in)")
	myPhone.setBodySim("Nano-SIM card & eSIM")
	myPhone.setDisplayType("Monochrome graphic")
	myPhone.setDisplay_Resolution("480 x 800 pixels, 5:3 ratio (~267 ppi density)")
	// Counting number of sensors
	myPhone.featuresSensors = "FaceID, Fingerprint (rear-mounted), accelerometer, proximity"
	var sensors = myPhone.countSensors()
	fmt.Printf("Number of sensors in my phone: %d\n", sensors)

	// Test 2
	// Showing that displaySize, bodyWeight, and launchAnnounced are the correct data types
	myPhone.displaySize = 3.5
	myPhone.bodyWeight = 162
	myPhone.launchAnnounced = 1999
	fmt.Printf("\nData Type of displaySize: %T\n", myPhone.displaySize)
	fmt.Printf("Data Type of bodyWeight: %T\n", myPhone.bodyWeight)
	fmt.Printf("Data Type of launchAnnounced: %T\n", myPhone.launchAnnounced)

	// Determining compatability
	var target = "Android 11"
	myPhone.setPlatformOs("Android 10")
	var compatibility = myPhone.isPlatformCompatible(target)
	fmt.Printf("\nOS Compatible: %t\n", compatibility)
	myPhone.platformOs = "Android 11"
	compatibility = myPhone.isPlatformCompatible(target)
	fmt.Printf("OS Compatible: %t\n", compatibility)
	if compatibility {
		fmt.Printf("My phone is compatible with %s\n", target)
	} else {
		fmt.Printf("My phone is not compatible with %s\n", target)
	}

	hasFaceID := myPhone.hasFaceID()

	if hasFaceID {
		fmt.Printf("\nMy phone has FaceID sensor\n")
	} else {
		fmt.Printf("\nMy phone does not have FaceID sensor\n")
	}

	// Determining weight class of my phone
	myPhone.bodyWeight = 120.5
	myPhone.weightClass()

	// Check if phone is available
	myPhone.setLaunchStatus("Available. Released 2020, May 11")
	isAvailable := myPhone.isAvailable()

	if isAvailable {
		fmt.Printf("\nMy phones launch status is available\n")
	} else {
		fmt.Printf("\nMy phones launch status is not available\n")
	}

	// Before reset
	myPhone.displayDetails()

	// Resetting phone
	myPhone.reset()

	// Display phones reset
	myPhone.displayDetails()

	// Test 3
	// Showing that the body dimension is set to null and body weight set to 0
	cellMap[2].displayDetails()

	// Report functions in use
	// Finding year with most releases
	year := yearWithMostLaunches(cellMap)

	if year > 0 {
		fmt.Printf("\nYear with the most phone launches (later than 1999): %d\n", year)
	} else {
		fmt.Println("\nNo launches found in years later than 1999.")
	}
	// Finding how many phones have one sensor
	count := oneSensor(cellMap)
	fmt.Printf("\nNumber of phones with only one feature sensor: %d\n", count)
	// Finding which phones released atleast one year later
	results := oneYearLater(cellMap)
	fmt.Println("\nPhones announced in one year and released in another year:")
	for _, phone := range results {
		fmt.Printf("%s, %s\n", phone.oem, phone.model)
	}
	// Finding the highest average body weight
	highestAvgOEM := findHighestAvgBodyWeight(cellMap)
	fmt.Printf("\nOEM with the highest average body weight: %s\n", highestAvgOEM)
}

// Report Functions
// Finds which year has the most phones launched
func yearWithMostLaunches(cells map[int]Cell) int {
	launchCount := make(map[int]int)

	for _, cell := range cells {
		year := cell.launchAnnounced

		if year > 1999 {
			launchCount[year]++
		}
	}

	maxYear := 0
	maxCount := 0

	for year, count := range launchCount {
		if count > maxCount {
			maxYear = year
			maxCount = count
		}
	}

	return maxYear
}

// Finds how many phones have only one sensor
func oneSensor(cells map[int]Cell) int {
	count := 0

	for _, cell := range cells {
		sensors := strings.Split(cell.featuresSensors, ",")

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

// Finds the OEM with highest average body weight
func findHighestAvgBodyWeight(cellMap map[int]Cell) string {
	weightSum := make(map[string]float64)
	count := make(map[string]int)

	for _, cell := range cellMap {
		weightSum[cell.oem] += cell.bodyWeight
		count[cell.oem]++
	}

	var highestAvgOEM string
	var maxAvgWeight float64

	for oem, sum := range weightSum {
		avgWeight := sum / float64(count[oem])
		if avgWeight > maxAvgWeight {
			maxAvgWeight = avgWeight
			highestAvgOEM = oem
		}
	}

	return highestAvgOEM
}

// Finds how many phones released atleast a year after announcement
func oneYearLater(cellMap map[int]Cell) []Cell {
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

// Extracts release year (for findDifferentYears function)
func extractReleaseYear(launchStatus string) int {
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

	return 0
}

// Takes the first four digits (for launch year)
func extractFirstFourDigits(input string) string {
	var digits []rune
	digitCount := 0

	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			digitCount++
			if digitCount == 4 {
				break
			}
		} else {
			digits = nil
			digitCount = 0
		}
	}

	return string(digits)
}

// Turns a string to an integer
func stringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// Turns a float to a string
func floatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

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
// Displays all the phones information
func (phone Cell) displayDetails() {
	fmt.Printf("\nOEM: %s\n", phone.oem)
	fmt.Printf("Model: %s\n", phone.model)
	fmt.Printf("Launch Announced: %d\n", phone.launchAnnounced)
	fmt.Printf("Launch Status: %s\n", phone.launchStatus)
	fmt.Printf("Body Dimensions: %s\n", phone.bodyDimensions)
	fmt.Printf("Body Weight: %.2f\n", phone.bodyWeight)
	fmt.Printf("Body SIM: %s\n", phone.bodySim)
	fmt.Printf("Display Type: %s\n", phone.displayType)
	fmt.Printf("Display Size: %.2f\n", phone.displaySize)
	fmt.Printf("Display Resolution: %s\n", phone.displayResolution)
	fmt.Printf("Features & Sensors: %s\n", phone.featuresSensors)
	fmt.Printf("Platform OS: %s\n", phone.platformOs)
}

// Resets everything to empty
func (phone *Cell) reset() {
	phone.oem = ""
	phone.model = ""
	phone.launchAnnounced = 0
	phone.launchStatus = ""
	phone.bodyDimensions = ""
	phone.bodyWeight = 0.0
	phone.bodySim = ""
	phone.displayType = ""
	phone.displaySize = 0.0
	phone.displayResolution = ""
	phone.featuresSensors = ""
	phone.platformOs = ""
}

// Gets the number of sensors that the phone has
func (phone *Cell) countSensors() int {
	sensorNames := strings.Split(phone.featuresSensors, ",")

	count := 0

	for _, sensor := range sensorNames {
		sensor = strings.TrimSpace(sensor)

		if sensor != "" {
			count++
		}
	}

	return count
}

// Checks if the phone is compatible with a certain OS
func (phone *Cell) isPlatformCompatible(targetOs string) bool {
	normalizedPlatformOs := strings.ToLower(strings.TrimSpace(phone.platformOs))
	normalizedTargetOs := strings.ToLower(strings.TrimSpace(targetOs))

	return strings.Contains(normalizedPlatformOs, normalizedTargetOs)
}

// Calculates the weight class of the phone
func (phone *Cell) weightClass() {
	const (
		lightWeightUpperLimit = 150.0
		heavyWeightLowerLimit = 200.0
	)

	if phone.bodyWeight <= lightWeightUpperLimit {
		fmt.Printf("The Phone Is Lightweight\n")
	} else if phone.bodyWeight > heavyWeightLowerLimit {
		fmt.Printf("\nThe Phone Is Heavy\n")
	} else {
		fmt.Printf("\nThe Phone Is Standard\n")
	}
}

// Check if the phone has faceID
func (phone *Cell) hasFaceID() bool {
	featuresLower := strings.ToLower(phone.featuresSensors)

	return strings.Contains(featuresLower, "faceid")
}

// Checks if the phone is available
func (phone *Cell) isAvailable() bool {
	statusLower := strings.ToLower(phone.launchStatus)

	return strings.Contains(statusLower, "available")
}

// Setter and getter functions for Cell
func (phone Cell) getOem() string {
	return phone.oem
}
func (phone *Cell) setOem(tempOem string) {
	phone.oem = tempOem
}
func (phone Cell) getModel() string {
	return phone.model
}
func (phone *Cell) setModel(tempModel string) {
	phone.model = tempModel
}
func (phone Cell) getLaunchAnnounced() int {
	return phone.launchAnnounced
}
func (phone *Cell) setLaunchAnnounced(tempLaunchAnnounced int) {
	phone.launchAnnounced = tempLaunchAnnounced
}
func (phone Cell) getLaunchStatus() string {
	return phone.launchStatus
}
func (phone *Cell) setLaunchStatus(tempLaunchStatus string) {
	phone.launchStatus = tempLaunchStatus
}
func (phone Cell) getBodyDimensions() string {
	return phone.bodyDimensions
}
func (phone *Cell) setBodyDimensions(tempbodyDimensions string) {
	phone.bodyDimensions = tempbodyDimensions
}
func (phone Cell) getBodyWeight() float64 {
	return phone.bodyWeight
}
func (phone *Cell) setBodyWeight(tempBodyWeight float64) {
	phone.bodyWeight = tempBodyWeight
}
func (phone Cell) getBodySim() string {
	return phone.bodySim
}
func (phone *Cell) setBodySim(tempBodySim string) {
	phone.bodySim = tempBodySim
}
func (phone Cell) getDisplayType() string {
	return phone.displayType
}
func (phone *Cell) setDisplayType(tempDisplayType string) {
	phone.displayType = tempDisplayType
}
func (phone Cell) getDisplaySize() float64 {
	return phone.displaySize
}
func (phone *Cell) setDisplaySize(tempDisplaySize float64) {
	phone.displaySize = tempDisplaySize
}
func (phone Cell) getDisplayResolution() string {
	return phone.displayResolution
}
func (phone *Cell) setDisplay_Resolution(tempDisplayResolution string) {
	phone.displayResolution = tempDisplayResolution
}
func (phone Cell) getFeaturesSensors() string {
	return phone.featuresSensors
}
func (phone *Cell) setFeaturesSensors(tempFeaturesSensors string) {
	phone.featuresSensors = tempFeaturesSensors
}
func (phone Cell) gePlatformOs() string {
	return phone.platformOs
}
func (phone *Cell) setPlatformOs(tempPlatformOs string) {
	phone.platformOs = tempPlatformOs
}
