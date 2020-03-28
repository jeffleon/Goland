package main

import "fmt"

func main() {
	// var nombre typo
	/* nombre no puede cambiar de tipo
	   := solo puede usarse para variables nuevas  */
	var numero int

	nombre := "Jefferson"
	nombre2:= "Juan"
	numero = 25
	numero += numero
	nombre, numero = "Pedro" , 70
	fmt.Println(numero)
	fmt.Println(nombre)
	fmt.Println(numero , nombre)
	nombre , nombre2 = nombre2 , nombre
	fmt.Println(nombre2 , nombre)
}
