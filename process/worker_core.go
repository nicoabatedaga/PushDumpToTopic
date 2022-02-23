package process

import (
	"encoding/csv"
	"fmt"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"os"
	"sync"
	"time"
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
	sizeChnReport = 50000
	sizeChnJob    = 50000
	cantOfWorkers = 200
	timeToScaling = 20
)

var pool *mpb.Progress
var wgpb sync.WaitGroup

func initializeWorkersData() {
	channOfReport = make(chan BAModel, sizeChnReport)
	jobs = make(chan Job, sizeChnJob)
	results = make(chan BAModel)
	response = []BAModel{}

	// Init pool of progress bar
	pool = mpb.New(mpb.WithWaitGroup(&wgpb))
	wgpb.Add(3)
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
	bar := pool.AddBar(int64(noOfWorkers),
		mpb.PrependDecorators(decor.Name("Workers running: ")),
		mpb.PrependDecorators(decor.CountersNoUnit("%v / %v")),
		mpb.PrependDecorators(decor.Percentage(decor.WCSyncSpace)),
		mpb.AppendDecorators(decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth), "done")),
	)
	for i := 0; i < noOfWorkers; i++ {
		bar.Increment()
		wg.Add(1)
		go worker(&wg)
		time.Sleep(time.Duration(int(time.Millisecond*timeToScaling) * i))
	}
	wg.Wait()
	wgpb.Done()
	close(channOfReport)
}

func allocate(ch []BAModel) {
	bar := pool.AddBar(int64(len(ch)),
		mpb.PrependDecorators(decor.Name("Jobs Created: ")),
		mpb.PrependDecorators(decor.CountersNoUnit("%v / %v")),
		mpb.PrependDecorators(decor.Percentage(decor.WCSyncSpace)),
		mpb.AppendDecorators(decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth), "done")),
	)
	for i := 0; i < len(ch); i++ {
		bar.Increment()
		job := Job{i, ch[i]}
		jobs <- job
	}
	fmt.Println("Close jobs channel")
	close(jobs)
	wgpb.Done()
}

func result(done chan bool, file *os.File, writer *csv.Writer) {

	defer file.Close()

	bar := pool.AddBar(0,
		mpb.PrependDecorators(decor.Name("Results write in file: ")),
		mpb.PrependDecorators(decor.CountersNoUnit("%v / %v")),
	)

	for report := range channOfReport {
		bar.Increment()
		if err := writer.Write([]string{report.BAID, report.Type, report.SiteID, report.UserID, report.Processed}); err != nil {
			fmt.Println(fmt.Sprintf("error al escribir %v", report.BAID))
		}
		writer.Flush()
		response = append(response, report)
		bar.SetTotal(int64(len(response)), false)
	}

	bar.SetTotal(int64(len(response)), true)
	wgpb.Done()
	wgpb.Wait()
	done <- true
}
