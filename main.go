package main

import (
	"bufio"
	"fmt"
	"github.com/mercadolibre/PushDumpToTopic/csv"
	"github.com/mercadolibre/PushDumpToTopic/process"
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
	listOfBugs := process.ReadCSV(datos_route)
	ts := time.Now()
	fmt.Println(fmt.Sprintf("Time start %v", ts))
	process.Analyze(listOfBugs, resultados_route)
	te := time.Now()
	fmt.Println(fmt.Sprintf("Time end %v", te))
	fmt.Println(fmt.Sprintf("Total time %v", te.Sub(ts).Seconds()))
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
