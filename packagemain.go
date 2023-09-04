// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	// Abre el archivo Excel
	excelFileName := "Validacion_Convocatoria_G57_29-08-23_DPEE.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error al abrir el archivo Excel: %v\n", err)
		os.Exit(1)
	}

	// Recorre todas las hojas del archivo Excel
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Leyendo hoja: %s\n", sheet.Name)

		// Recorre todas las filas de la hoja
		for _, row := range sheet.Rows {
			// Recorre todas las celdas de la fila
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}
}
