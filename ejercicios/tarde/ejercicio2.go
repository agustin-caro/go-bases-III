package tarde

import "fmt"

type Producto struct {
	id       float32
	Nombre   string
	Precio   float32
	cantidad float32
}

type UsuarioCommerce struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto //array de puntero a productos
}

func Ejercicio2() {

	p1 := Producto{0, "azucar", 10, 1}
	p2 := Producto{1, "yerba", 20, 5}
	p3 := Producto{2, "asado", 35, 4}

	var Productos []Producto

	Productos = append(Productos, p1)
	Productos = append(Productos, p2)
	Productos = append(Productos, p3)

	var Usuarios []*UsuarioCommerce

	var ProductosPTR []*Producto
	ProductosPTR = append(ProductosPTR, &Productos[0])
	fmt.Printf("Length: %v, capacity: %v \n", len(ProductosPTR), cap(ProductosPTR))

	u1 := UsuarioCommerce{
		Nombre:    "Pablo",
		Apellido:  "Garcia",
		Correo:    "test",
		Productos: ProductosPTR,
	}

	ProductosPTR = nil
	ProductosPTR = append(ProductosPTR, &Productos[1])
	ProductosPTR = append(ProductosPTR, &Productos[2])

	//Usuario 2
	u2 := UsuarioCommerce{
		Nombre:    "Gabriel",
		Apellido:  "Lozada",
		Correo:    "Test",
		Productos: ProductosPTR,
	}

	Usuarios = append(Usuarios, &u1)
	Usuarios = append(Usuarios, &u2)

	for _, u := range Usuarios {
		fmt.Printf("Usuario %v: \n", u.Nombre)
		for _, v := range u.Productos {
			fmt.Printf(" Producto %v \t", v.Nombre)
		}
		fmt.Printf("\n-------\n")
	}

	//Nuevo producto
	p4 := nuevoProducto("Aceite", 50)
	Productos = append(Productos, p4)

	//agregar producto
	agregarProducto(&u1, &p4, 1)
	//eliminar productos
	borrarProducto(&u2)

	for _, u := range Usuarios {
		fmt.Printf("Usuario %v: \n", u.Nombre)
		for _, v := range u.Productos {
			fmt.Printf(" Producto %v \t", v.Nombre)
		}
		fmt.Printf("\n-------\n")
	}

}

func nuevoProducto(nombre string, precio float32) Producto {
	prod := Producto{
		Nombre: nombre,
		Precio: precio,
	}

	return prod
}

func agregarProducto(uPtr *UsuarioCommerce, pPtr *Producto, cantidad int) {

	uPtr.Productos = append(uPtr.Productos, pPtr)

}

func borrarProducto(uPtr *UsuarioCommerce) {
	uPtr.Productos = nil
	uPtr.Productos = nil
}
