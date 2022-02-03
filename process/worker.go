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

func Analyze(listOfBA []BAModel, resultados_route string) {

	channOfReport = make(chan BAModel, 50)
	jobs = make(chan Job, 50)
	results = make(chan BAModel)
	response = []BAModel{}

	go allocate(listOfBA)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 200
	createWorkerPool(noOfWorkers)
	//Todos los datos procesados
	<-done

	//En este punto tenemos los workers finalizados y con los resultados en el channel de report
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

func generateReport(bm BAModel) *BAModel {
	report := BAModel{
		BAID:      bm.BAID,
		Processed: "0",
	}
	/*
		user, err := services.GetUser(bm.UserID)
		//fmt.Println(fmt.Sprintf("GET de user %v, con error :%v", bm.UserID, err))
		if err != nil {
			report.UserIdentificationID = "error"
			report.UserCompanyIdentificationID = "error"

		} else {
			report.UserIdentificationID = user.Identification.Number
			report.UserCompanyIdentificationID = user.Company.Identification
			if strings.Compare(report.WithIdentificationID, report.UserIdentificationID) == 0 || strings.Compare(report.WithIdentificationID, report.UserCompanyIdentificationID) == 0 {
				report.AreEquals = 1
			} else {
				var x1, x2 float64
				x1, x2 = 0, 0
				if report.UserIdentificationID != "" {
					x1 = simhash.GetLikenessValue(report.WithIdentificationID, report.UserIdentificationID)
				}
				if report.UserCompanyIdentificationID != "" {
					x2 = simhash.GetLikenessValue(report.WithIdentificationID, report.UserCompanyIdentificationID)
				}
				if x1 > x2 {
					report.AreEquals = x1
				} else {
					report.AreEquals = x2
				}
			}
		}
	*/
	return &report
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
	for i := 0; i < len(ch); i++ {
		job := Job{i, ch[i]}
		jobs <- job
	}
	//el canal de job esta cargado completo
	close(jobs)
}

func result(done chan bool) {
	for report := range channOfReport {
		response = append(response, report)
	}
	done <- true
}
