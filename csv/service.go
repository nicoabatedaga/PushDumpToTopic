package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SplitCsv(route string, sizeOfFile int) {
	csvFile, err := os.Open(route)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	var listToCSV [][]string
	i := 0
	for _, line := range csvLines {
		listToCSV = append(listToCSV, line)
		if len(listToCSV) == sizeOfFile {
			newCsvFile, err := os.Create(fmt.Sprintf("%v.%v.csv", route, i))
			if err != nil {
				fmt.Println(err)
			}
			writer := csv.NewWriter(newCsvFile)
			writer.WriteAll(listToCSV)
			writer.Flush()
			newCsvFile.Close()
			listToCSV = [][]string{}
			i++
		}
	}
	if len(listToCSV) > 0 {
		newCsvFile, err := os.Create(fmt.Sprintf("%v.%v.csv", route, i))
		if err != nil {
			fmt.Println(err)
		}
		writer := csv.NewWriter(newCsvFile)
		writer.WriteAll(listToCSV)
		writer.Flush()
		newCsvFile.Close()
		listToCSV = [][]string{}
		i++
	}

}
