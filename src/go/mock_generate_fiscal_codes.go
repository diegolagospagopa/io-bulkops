package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var listCodiciFiscali []string

	for i := 0; i < 10; i++ {
		listCodiciFiscali = append(listCodiciFiscali, generateMockCodiceFiscale(i))
	}
	writeCodiciFiscaleInCsv(listCodiciFiscali)

}

func writeCodiciFiscaleInCsv(listCodiciFiscali []string) {
	fileCsv := generateCsvFile("mock_cf.csv")
	csvWriter := csv.NewWriter(fileCsv)

	for _, empRow := range listCodiciFiscali {
		fmt.Println(empRow)
		_ = csvWriter.Write(strings.Fields(empRow))
	}

	csvWriter.Flush()
	fileCsv.Close()
}

func generateCsvFile(filename string) *os.File {
	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	return csvFile
}

func generateMockCodiceFiscale(randomNum int) string {

	var rawHash string = getMD5Hash(strconv.Itoa(randomNum))
	codiceFiscaleMock := rawHash[:16]

	return codiceFiscaleMock
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
