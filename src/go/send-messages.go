package main

import (
	// "container/list"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	channel := make(chan string)
	var cfList [][]string = readCsvData("mock_cf.csv")

	go func (ch chan string)  {
		defer func() {
			fmt.Println("📢 call close")
			close(ch)
		}()
		for i, cf := range cfList {
			sendCode(cf[0], ch)
			fmt.Println("cf number ", i, " sended")
		}
	}(channel)


	for c := range channel {
		go func (item string)  {
			fmt.Println("Channel response: ", item)
		}(c)
	}
	fmt.Println("🚀 completed with success")
}

func sendCode(code string, channel chan string) {
	fmt.Println(code)

	var requestUrl = fmt.Sprintf("http://localhost:3000/%s", code)

	_, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(requestUrl, "might be down!")
		channel <- requestUrl
		return
	}

	fmt.Println(requestUrl, "is up!")
	channel <- requestUrl
}

func readCsvData(fileName string) [][]string {
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

	// fmt.Println(data)
	return data
}
