package process

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Analyze(listOfBA []BAModel, resultados_route string) {

	file, writer := openCSV(resultados_route)

	initializeWorkersData()

	go allocate(listOfBA)
	done := make(chan bool)
	go result(done, file, writer)
	createWorkerPool(cantOfWorkers)

	<-done //Todos los datos procesados

	fmt.Println(fmt.Sprintf("En result tenemos: %v ", len(response)))
}

func openCSV(resultados_route string) (*os.File, *csv.Writer) {
	csvFile, err := os.Create(resultados_route)
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(csvFile)

	return csvFile, w
}

// generateReport
func generateReport(bm BAModel) *BAModel {
	report := BAModel{
		BAID:      bm.BAID,
		Type:      bm.Type,
		SiteID:    bm.SiteID,
		UserID:    bm.UserID,
		Processed: bm.Processed,
	}
	if report.Processed != "1" {
		report.Processed = "0"
		/*if err := services.PostMsg(report.BAID, report.Type, report.SiteID, report.UserID); err == nil {
			report.Processed = "1"
		} else {
			//fmt.Println(err.Error())
		}*/
	}
	return &report
}
