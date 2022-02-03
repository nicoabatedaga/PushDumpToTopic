package main

import (
	"github.com/mercadolibre/PushDumpToTopic/process"
	"testing"
)

func Test1(t *testing.T) {
	listOfBugs := process.ReadCSV("/Users/nabatedaga/Desktop/datatrunc.csv")
	process.Analyze(listOfBugs, "/Users/nabatedaga/Desktop/datatrunc.csv.resultados.csv")
}
