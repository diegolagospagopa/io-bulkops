package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	readCsvData("mock_cf.csv")
}

func sendCode(code string)  {
	fmt.Println(code)
}

func readCsvData(fileName string) {
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
