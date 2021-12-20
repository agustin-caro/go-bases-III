package maniana

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
)

func Ejercicio2() {

	// var ProductosLeidos []Producto
	// fmt.Printf("ID \t\t Precio \t\t Cantidad\n")

	f, err := os.OpenFile("productos.csv", os.O_RDONLY, 0644)
	checkError(err)

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	checkError(err)

	w := tabwriter.NewWriter(os.Stdout, 10, 2, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintf(w, "ID \tPrecio\tCantidad\t\n")

	for _, record := range records {

		_, err := fmt.Fprintf(w, "%v \t %v \t %v\t\n", record[0], record[1], record[2])
		if err != nil {
			panic(err)
		}

		w.Flush()
		// fmt.Printf(fmt.Sprintf("%v \t\t %v \t\t %v\b", record[0], record[1], record[2]))
	}

}
