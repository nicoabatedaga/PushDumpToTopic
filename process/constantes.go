package process

import (
	"encoding/csv"
	"fmt"
	"os"
)

type BAModel struct {
	BAID      string `json:"id"`
	Type      string `json:"type"`
	SiteID    string `json:"site_id"`
	UserID    string `json:"user_id"`
	Processed string `json:"processed"`
}

func ReadCSV(datos_route string) []BAModel {
	csvFile, err := os.Open(datos_route)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	listOfBA := []BAModel{}
	for _, line := range csvLines {
		bm := BAModel{
			BAID:      line[0],
			Type:      line[1],
			SiteID:    line[2],
			UserID:    line[3],
			Processed: "0",
		}
		if len(line) > 4 {
			bm.Processed = line[4]
		}
		listOfBA = append(listOfBA, bm)
		//fmt.Println(bm.WithID + " " + bm.UserID + " " + bm.WithIdentificationID)
	}
	return listOfBA
}
