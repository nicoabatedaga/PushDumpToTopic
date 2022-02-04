package process

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type Job struct {
	id int
	BA BAModel
}

//Armo las variables para los workers
var channOfReport chan BAModel
var jobs chan Job
var results chan BAModel
var response []BAModel

const (
	sizeChnReport = 200
	sizeChnJob    = 200
	cantOfWorkers = 100
)

func initializeWorkersData() {
	channOfReport = make(chan BAModel, sizeChnReport)
	jobs = make(chan Job, sizeChnJob)
	results = make(chan BAModel)
	response = []BAModel{}
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := generateReport(job.BA)
		//fmt.Println(fmt.Sprintf("Job: %v - Output para el generate report: %#v", job.id, *output))
		//fmt.Println("job terminado:", job.id)
		channOfReport <- *output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(channOfReport)
}

func allocate(ch []BAModel) {
	fmt.Println(fmt.Sprintf("Tenemos %v datos", len(ch)))
	for i := 0; i < len(ch); i++ {
		job := Job{i, ch[i]}
		jobs <- job
	}
	//el canal de job esta cargado completo
	close(jobs)
}

func result(done chan bool, file *os.File, writer *csv.Writer) {
	for report := range channOfReport {
		//fmt.Println(fmt.Sprintf("vamos a escribir: %v", report.BAID))
		if err := writer.Write([]string{report.BAID, report.Processed}); err != nil {
			fmt.Println(fmt.Sprintf("error al escribir %v", report.BAID))
		}
		writer.Flush()
		response = append(response, report)
	}
	file.Close()
	done <- true
}
