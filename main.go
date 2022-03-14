package main

import (
	"bufio"
	"fmt"
	"github.com/nicoabatedaga/PushDumpToTopic/csv"
	"github.com/nicoabatedaga/PushDumpToTopic/process"
	"os"
	"strings"
	"time"
)

func main() {
	processCSV()
}

func splitCSV() {
	datos_route := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ruta complea del archivo de datos: ")
	scanner.Scan()
	datos_route = scanner.Text()
	ts := time.Now()
	fmt.Println(fmt.Sprintf("Time start %v", ts))
	csv.SplitCsv(datos_route, 5000000)
	te := time.Now()
	fmt.Println(fmt.Sprintf("Time end %v", te))
	fmt.Println(fmt.Sprintf("Total time %v", te.Sub(ts).Seconds()))
}

func printFailedResults() {
	datos_route := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ruta complea del archivo de datos: ")
	scanner.Scan()
	datos_route = scanner.Text()
	ts := time.Now()
	fmt.Println(fmt.Sprintf("Time start %v", ts))
	csv.AnalizeResponse(datos_route)
	te := time.Now()
	fmt.Println(fmt.Sprintf("Time end %v", te))
	fmt.Println(fmt.Sprintf("Total time %v", te.Sub(ts).Seconds()))
}

func processCSV() {
	datos_route := ""
	resultados_route := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ruta complea del archivo de datos: ")
	scanner.Scan()
	datos_route = scanner.Text()
	baseResultado := strings.Replace(datos_route, ".csv", "", -1)
	resultados_route = baseResultado + ".resultados.csv"
	ts := time.Now()
	fmt.Println(fmt.Sprintf("Time start %v", ts))
	processFiles(datos_route, resultados_route)
	csv.AnalizeResponse(resultados_route)
	fmt.Println("resultados_route:", resultados_route)
	te := time.Now()
	fmt.Println(fmt.Sprintf("Time end %v", te))
	fmt.Println(fmt.Sprintf("Total time %v", te.Sub(ts).Seconds()))
}

func processFiles(datos_route, resultados_route string) int {
	listOfBugs := process.ReadCSV(datos_route)
	process.Analyze(listOfBugs, resultados_route)
	return csv.AnalizeResponse(resultados_route)
}

func mergeCSV() {
	datos_base := ""
	datos_incompleto := ""
	resultados_route := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ruta complea del archivo de datos base: ")
	scanner.Scan()
	datos_base = scanner.Text()
	fmt.Print("ruta complea del archivo de datos incompletos: ")
	scanner.Scan()
	datos_incompleto = scanner.Text()
	baseResultado := strings.Replace(datos_base, ".csv", "", -1)
	resultados_route = baseResultado + ".merge.csv"
	listOfBACompleta := process.ReadCSV(datos_base)
	listOfBAIncompleta := process.ReadCSV(datos_incompleto)
	ts := time.Now()
	fmt.Println(fmt.Sprintf("Time start %v", ts))
	csv.MergeCSV(listOfBAIncompleta, listOfBACompleta, resultados_route)
	te := time.Now()
	fmt.Println(fmt.Sprintf("Time end %v", te))
	fmt.Println(fmt.Sprintf("Total time %v", te.Sub(ts).Seconds()))
}
