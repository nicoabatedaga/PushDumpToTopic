package main

import (
	"bufio"
	"fmt"
	"github.com/mercadolibre/PushDumpToTopic/process"
	"os"
)

func main() {
	datos_route:=""
	resultados_route:=""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ruta complea del archivo de datos: ")
	scanner.Scan()
	datos_route = scanner.Text()
	resultados_route = datos_route + ".resultados.csv"
	listOfBugs := process.ReadCSV(datos_route)
	process.Analyze(listOfBugs,resultados_route)
}
