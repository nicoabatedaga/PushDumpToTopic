package main

import (
	"github.com/mercadolibre/PushDumpToTopic/process"
	"testing"
)

func Test1(t *testing.T) {
	listOfBugs := process.ReadCSV("/Users/nabatedaga/Desktop/data.csv")
	process.Analyze(listOfBugs, "/Users/nabatedaga/Desktop/data.csv.resultados.csv")
}
