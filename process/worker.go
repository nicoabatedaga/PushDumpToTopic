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
	//En este punto tenemos los workers finalizados y con los resultados en el channel de report
	//fillCSV(resultados_route)
}

func openCSV(resultados_route string) (*os.File, *csv.Writer) {
	fmt.Println(fmt.Sprintf("vamos a abrir el archivo"))
	csvFile, err := os.Create(resultados_route)
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(csvFile)
	headers := []string{"ba_id", "processed"}
	fmt.Println(fmt.Sprintf("vamos a escribir los headers"))
	if errheader := w.Write(headers); errheader != nil {
		fmt.Println("error en el write de headers")
	}
	fmt.Println(fmt.Sprintf("retorno file y writer"))
	return csvFile, w
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
