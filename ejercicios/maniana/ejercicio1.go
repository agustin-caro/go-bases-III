package maniana

import (
	"fmt"
	"os"
)

type Producto struct {
	id       float32
	precio   float32
	cantidad float32
}

func Ejercicio1() {

	p1 := Producto{
		id:       0,
		precio:   10,
		cantidad: 2,
	}

	p2 := Producto{
		id:       10,
		precio:   20,
		cantidad: 6,
	}

	var Productos []Producto
	Productos = append(Productos, p1)
	Productos = append(Productos, p2)

	//si no existe, crea el archivo. Luego, hace append
	f, err := os.OpenFile("productos.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)

	for _, p := range Productos {

		strout := fmt.Sprintf(" %v,%v,%v \n", p.id, p.precio, p.cantidad)
		_, err = f.WriteString(strout)
		checkError(err)

	}

	f.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error al escribir archivo")
		panic(err)
	}
}
