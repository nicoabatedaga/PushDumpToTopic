package main

import (
	"github.com/mercadolibre/PushDumpToTopic/process"
	"testing"
)

func TestDataTrunc(t *testing.T) {
	listOfBugs := process.ReadCSV("/Users/nabatedaga/Desktop/datatrunc.csv")
	process.Analyze(listOfBugs, "/Users/nabatedaga/Desktop/datatrunc.csv.resultados.csv")
}

func TestData(t *testing.T) {
	listOfBugs := process.ReadCSV("/Users/nabatedaga/Desktop/data.csv")
	process.Analyze(listOfBugs, "/Users/nabatedaga/Desktop/data.csv.resultados.csv")
}
