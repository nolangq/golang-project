package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Cell struct {
	oem               string
	model             string
	launchAnnounced   string
	launchStatus      string
	bodyDimensions    string
	bodyWeight        string
	bodySim           string
	displayType       string
	displaySize       string
	displayResolution string
	featuresSensors   string
	platformOs        string
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
func (phone Cell) getLaunchAnnounced() string {
	return phone.launchAnnounced
}
func (phone Cell) setLaunchAnnounced(tempLaunchAnnounced string) {
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
func (phone Cell) getBodyWeight() string {
	return phone.bodyWeight
}
func (phone Cell) setBodyWeight(tempBodyWeight string) {
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
func (phone Cell) getDisplaySize() string {
	return phone.displaySize
}
func (phone Cell) setDisplaySize(tempDisplaySize string) {
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

/*
func (phone *Cell) DeleteObject() {
	var newPhone = &Cell{}
	*phone = *newPhone
}
*/

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

	// Process each row
	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}

	// Testing struct methods
	/*
		myPhone := Cell{}
		myPhone.model = "iPhone"
		fmt.Println(myPhone.model)
		myPhone.DeleteObject()
		fmt.Println(myPhone.model)
	*/
}
