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
