package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/nicoabatedaga/PushDumpToTopic/process"
	"io"
	"os"
	"strings"
)

func SplitCsv(route string, sizeOfFile int) {
	csvFile, err := os.Open(route)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")

	reader := csv.NewReader(csvFile)
	//if err != nil {
	//	fmt.Println("ERROR:", err)
	//}
	var listToCSV [][]string
	baseSalida := strings.Replace(route, ".csv", "", -1)
	i := 0
	numberOfLine := 0
	line, err := reader.Read()
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR:%v, numero de linea:%v", err, numberOfLine))
	}
	for line != nil {
		listToCSV = append(listToCSV, line)
		if len(listToCSV) == sizeOfFile {
			newCsvFile, err := os.Create(fmt.Sprintf("%v%v.csv", baseSalida, i))
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
		line, err = reader.Read()
		if err != nil {
			fmt.Println(fmt.Sprintf("ERROR:%v, numero de linea:%v", err, numberOfLine))
		}
		numberOfLine++
	}
	if len(listToCSV) > 0 {
		newCsvFile, err := os.Create(fmt.Sprintf("%v%v.csv", baseSalida, i))
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

func AnalizeResponse(route string) int {
	csvFile, err := os.Open(route)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")

	reader := csv.NewReader(csvFile)
	listOfBA := []process.BAModel{}
	numberOfLine := 0
	countOfError := 0
	line, err := reader.Read()
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR:%v, numero de linea:%v", err, numberOfLine))
	}
	for line != nil {
		bm := process.BAModel{
			BAID:      line[0],
			Type:      line[1],
			SiteID:    line[2],
			UserID:    line[3],
			Processed: line[4],
		}
		listOfBA = append(listOfBA, bm)
		if bm.Processed == "0" {
			countOfError++
			//fmt.Println(fmt.Sprintf("ERROR line:%v - BA.id:%v", numberOfLine, bm.BAID))
		}
		line, err = reader.Read()
		if err != nil && err != io.EOF {
			//fmt.Println(fmt.Sprintf("ERROR:%v, numero de linea:%v", err, numberOfLine))
		}
		numberOfLine++
	}

	fmt.Println(fmt.Sprintf("Cantidad de errores en el file: %v", countOfError))

	return countOfError
}

func MergeCSV(incompleto, base []process.BAModel, resultado_route string) {
	newCsvFile, err := os.Create(resultado_route)
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(newCsvFile)
	var mapOfIdsIncompleted map[string]process.BAModel = make(map[string]process.BAModel)
	for _, incompletoItem := range incompleto {
		if err := writer.Write([]string{incompletoItem.BAID, incompletoItem.Type, incompletoItem.SiteID, incompletoItem.UserID, incompletoItem.Processed}); err != nil {
			fmt.Println(fmt.Sprintf("error al escribir %v", incompletoItem.BAID))
		} else {
			mapOfIdsIncompleted[incompletoItem.BAID] = incompletoItem
		}
		writer.Flush()
	}
	for _, baseItem := range base {
		if _, ok := mapOfIdsIncompleted[baseItem.BAID]; !ok {
			if err := writer.Write([]string{baseItem.BAID, baseItem.Type, baseItem.SiteID, baseItem.UserID, baseItem.Processed}); err != nil {
				fmt.Println(fmt.Sprintf("error al escribir %v", baseItem.BAID))
			}
		}
		writer.Flush()
	}
	newCsvFile.Close()
}
