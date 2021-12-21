package tarde

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       string
	Correo     string
	Contrasena string
}

func Ejercicio1() {

	u := Usuario{
		Nombre:     "Agustin",
		Apellido:   "Caro",
		Edad:       "40",
		Correo:     "test",
		Contrasena: "pwd",
	}

	fmt.Printf("%v\n", u)

	u.cambiarNombreApellido("Carlos", "Gardel")
	fmt.Printf("%v", u)

}

// Que significa lo siguiente
func (u *Usuario) cambiarNombreApellido(nombreNuevo string, apeNuevo string) {
	u.Nombre = nombreNuevo
	u.Apellido = apeNuevo
}

func (u *Usuario) cambiarEdad(edad string) {
	u.Edad = edad
}
func (u *Usuario) cambiarContrasena(contrasena string) {
	u.Contrasena = contrasena
}
func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

//
