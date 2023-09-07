package main

import (
	"fmt"
	"os"

	"regexp"

	"github.com/go-gota/gota/dataframe"
	"github.com/tealeg/xlsx"
)

var Cantidad_grupos int

//	func Grupos(x, y int) string {
//		return "G", x, "M", y, "-R-001"
//	}
func funcionPersonalizada(row2 []string) []interface{} {
	resultados := make([]interface{}, len(row2))

	for i, valor := range row2 {
		resultados[i] = valor
	}

	return resultados
}
func main() {
	// Ruta al archivo Excel que deseas leer
	excelFile := "Asignaciones_G27M23-R.xlsx"

	// Abrir el archivo Excel
	xlFile, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Println("Error al abrir el archivo Excel:", err)
		os.Exit(1)
	}

	// Crear un DataFrame vacío

	df := dataframe.New()

	// Crear un DataFrame concatenando las Series
	// Recorrer las hojas del archivo Excel
	for _, sheet := range xlFile.Sheets {
		// Crear un DataFrame para cada hoja
		data := [][]string{}
		for _, row := range sheet.Rows {
			rowValues := []string{}
			for _, cell := range row.Cells {
				rowValues = append(rowValues, cell.String())

			}
			data = append(data, rowValues)

		}

		// Crear un DataFrame para la hoja actual
		sheetDf := dataframe.LoadRecords(data)
		//fmt.Println(sheetDf)
		// Agregar el DataFrame de la hoja actual al DataFrame principal
		cadena := "Alice,25;Bob,30;Charlie,22"
		substring := cadena[0:5] // Extrae los primeros 5 caracteres
		fmt.Println(substring)
		// Dividir la cadena en elementos utilizando una expresión regular
		elementos := regexp.MustCompile("[;,]").Split(cadena, -1)
		fmt.Println(elementos)
		df := sheetDf
		Nf := df.Nrow()
		fmt.Println(df, Nf)
		folios := df.Col("Folio Alumno")
		fmt.Println(folios)
		// Iterar a través de las filas y aplicar la función personalizada
		for _, row3 := range df.Records() {
			resultado := funcionPersonalizada(row3)
			//fmt.Println(resultado)
		}
	}
	// Imprimir el DataFrame
	fmt.Println(df)
}
