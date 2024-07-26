// Go Bases Tarea 1 - Dificultad: 2.5/5
// Dentro de nuestro codigo vamos a tener 2 variables definidas, que van a ser:

// license: si la persona tiene licencia
// age: la edad de la persona
//   Si la persona tiene mas de 15 años y tiene licencia, debemos imprimir un mensaje que diga "Puede seguir avanzando"
//   Si la persona tiene 15 años o menos, o no tiene licencia, debemos imprimir un mensaje que diga "No puede seguir circulando"

// package main
// import "fmt"

// func main() {
//     license := true
//     age := 19
//     // tarea a realizar
// }

// Por ejemplo:
// si license es true y age es 19 debe mostrar el siguiente mensaje: Puede seguir avanzando
// si license es false y age es 19 debe mostrar el siguiente mensaje: No puede seguir circulando
// si license es true y age es 15 debe mostrar el siguiente mensaje: No puede seguir circulando

package main

import (
	"fmt"
)

func main() {

	var license bool = false
	var age int = 18

	switch {
	case !license || age < 18:
		fmt.Println("No Puede continuar")
	case license && age >= 18:
		fmt.Println("Puede continuar")
	default:
		fmt.Println("No valido")
	}
}
