package process

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Analyze(listOfBA []BAModel, resultados_route string) {

	initializeWorkersData()

	go allocate(listOfBA)
	done := make(chan bool)
	go result(done)
	createWorkerPool(cantOfWorkers)
	<-done //Todos los datos procesados

	//En este punto tenemos los workers finalizados y con los resultados en el channel de report
	fillCSV(resultados_route)
}

func fillCSV(resultados_route string) {
	csvFile, err := os.Create(resultados_route)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(csvFile)
	headers := []string{"ba_id", "processed"}
	var list [][]string
	for _, x := range response {
		list = append(list, []string{x.BAID, x.Processed})
	}
	//fmt.Println(list)

	if errheader := w.Write(headers); errheader != nil {
		fmt.Println("error en el write de headers")
	}
	if errlist := w.WriteAll(list); errlist != nil {
		fmt.Println("error en el write de values")
	}
}

// generateReport
func generateReport(bm BAModel) *BAModel {
	report := BAModel{
		BAID:      bm.BAID,
		Processed: bm.Processed,
	}
	return &report
}
