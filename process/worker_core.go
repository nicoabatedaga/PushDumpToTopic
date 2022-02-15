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

var channOfReport chan BAModel
var jobs chan Job
var results chan BAModel
var response []BAModel

const (
	sizeChnReport = 200
	sizeChnJob    = 200
	cantOfWorkers = 1
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
	for i := 0; i < len(ch); i++ {
		job := Job{i, ch[i]}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool, file *os.File, writer *csv.Writer) {

	defer file.Close()

	for report := range channOfReport {
		if err := writer.Write([]string{report.BAID, report.Type, report.SiteID, report.UserID, report.Processed}); err != nil {
			fmt.Println(fmt.Sprintf("error al escribir %v", report.BAID))
		}
		writer.Flush()
		response = append(response, report)
	}

	done <- true
}
