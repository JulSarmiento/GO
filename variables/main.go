package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// int: Cualquier tipo de valor
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	// Uint: Valores solo positivos
	var myUintVar uint
	myUintVar = 12
	fmt.Println("My Uint variable is: ", myUintVar)

	// string: Valores tipo string
	var myStringVar string
	myStringVar = "Soy un string"
	fmt.Println("My string variable is: ", myStringVar)

	// bool: Valores tipo booleano
	var myBooleanVar bool
	myBooleanVar = true
	fmt.Println("My boolean variable is: ", myBooleanVar)

	// usamos un "&" para indicar la direccion de la memoria en la que se esta guardando la variable
	fmt.Println("My variable address is: ", &myStringVar)

	// Podemos guardar un dato sin asignarla a la variable, empleando ":" antes del signo igual.
	// Asigan el tipo de dato segun lo que contenga
	myIntVar2 := 12
	fmt.Println("My variable 2 is: ", myIntVar2)

	// Las constantes son inmutables y toman el tipo de dato del valor asignado.
	const myConst = 2
	fmt.Println("My const is: ", myConst)

	// Tambien podemos indicar el tipo de dato.
	const myIntconst int = 12
	fmt.Println("My const is: ", myIntconst)

	// Tambien podemos indicar el tipo de dato.
	const myStringconst string = "holis"
	fmt.Println("My const is: ", myStringconst)

	fmt.Println()

	var my8BitsIntVar int8 // si no declaro valor, asigna por defecto 0

	// Printf: Imprime pendiente un formato puntual
	// %d (i, para entero, s para string) hace referencia a valor numerico
	// \n hace referencia salto de linea
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16 // si no declaro valor, asigna por defecto 0
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32 // si no declaro valor, asigna por defecto 0
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64 // si no declaro valor, asigna por defecto 0
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	var myDefaulBitsIntVar int // si no declaro valor, asigna por defecto 0
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myDefaulBitsIntVar, unsafe.Sizeof(myDefaulBitsIntVar), unsafe.Sizeof(myDefaulBitsIntVar)*8)

	// La comilla francesa permite el uso de de multilienas
	myString2 := `
		My String variable in goland
		with
		multiple
		line
	`

	fmt.Printf("The variable value is: %s\n ", myString2)

	// Scope
	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%f", floatVar) // Sprintf nos retorna un string
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar) // Sprintf nos retorna un string
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err) // retorna nill, que representa el null
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa1333", 0, 64) // valor errado que termina setenado la variable en 0 con mensaje de error
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		// si no quiero usar el error, puedo emplear un guion bajo

		floatVar1, _ := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("Type: %T, value: %f\n", floatVar1, floatVar1)

	}
}
