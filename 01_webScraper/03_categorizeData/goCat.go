package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// function to check for errors
func checkErr(task string, err error) {
	if err != nil {
		log.Fatalf("%s Failure,error : %s", task, err)
		return
	}
}

func main() {

	// read csv into records
	records := readCsvFile("data_C.csv")

	// create a new file named
	data, dataErr := os.Create("data_D.csv")
	// check for errors
	checkErr("File Creation", dataErr)
	// close file afterwards
	defer data.Close()

	// instantiate a csv writer
	writer := csv.NewWriter(data)
	// flush contents afterwards
	defer writer.Flush()

	// string to check already visited profiles
	visitedlist := ""
	// int to check number of respective profiles ~ count of app dev profiles
	count := 0
	// stipend as string
	aStipend := ""
	// stipend as integer for calculation
	iStipend := 0
	// temp file for str to int conversion
	temp := 0

	// looping through all profiles one bye one
	for i := 0; i < len(records); i++ {

		// reset count to 0 for each new profile
		count = 0
		// reset stipend to 0 for each new profile
		iStipend = 0

		// if the profile is already visited
		if strings.Contains(visitedlist, records[i][0]) == true {
			// skip the profile
			continue
		}
		// otherwise add the profile to visited list
		visitedlist += string(records[i][0])
		// and loop through the database looking for current profile
		for j := 0; j < len(records); j++ {

			// if current profile is found
			if records[i][0] == records[j][0] {

				// increment count
				count++

				// and add its stipend to currrent profile stipend
				aStipend = string(records[j][1])
				aStipend = strings.Replace(aStipend, " ", "", -1)
				temp, _ = strconv.Atoi(aStipend)
				iStipend += temp

			}

		}
		// when no more profiles natch current profile
		// write current profile, count & stipend into csv file
		writer.Write(
			[]string{records[i][0],
				strconv.Itoa(count), strconv.Itoa(iStipend)})

	}

	// let yourself know when you're done :D
	fmt.Printf("Job Complete !\n")

}

// function to read csv file
func readCsvFile(filePath string) [][]string {

	// open given file
	f, err := os.Open(filePath)
	// check for errors
	checkErr("Unable to read input file "+filePath, err)
	// close file afterwards
	defer f.Close()

	// instantiate a csv reader
	csvReader := csv.NewReader(f)
	// read files
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
