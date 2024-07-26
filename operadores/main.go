package main

import (
	"fmt"
)

func main() {
	var yearsOld uint = 32

	fmt.Println("Operadores")
	fmt.Println(yearsOld > 30)  // true
	fmt.Println(yearsOld < 32)  // false
	fmt.Println(yearsOld <= 32) // true
	fmt.Println(yearsOld >= 40) // false
	fmt.Println(yearsOld == 32) // true

	fmt.Println("Operadores logicos")

	// Operador OR
	fmt.Println("Operador OR")
	fmt.Println(yearsOld < 32 || yearsOld == 32) // true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // false
	fmt.Println(yearsOld < 40 || yearsOld == 33) // true

	// Operador AND
	fmt.Println("Operador AND")
	fmt.Println(yearsOld < 32 && yearsOld == 32) // true
	fmt.Println(yearsOld < 32 && yearsOld == 33) // false
	fmt.Println(yearsOld < 40 && yearsOld == 32) // true

	// Operador de negacion
	fmt.Println("Operador de negacion")
	fmt.Println(true)
	fmt.Println(!true)

	fmt.Println(yearsOld < 40)
	fmt.Println(!(yearsOld < 40))

	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)   // ((false and true) or true) = true
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40)) // (false and (true or true)) = false

	// Condicionales

	fmt.Println("Condicionales")

	yearsOld = 19

	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	var boolVar bool = true

	if boolVar {
		fmt.Println("is true!")
	} else {
		fmt.Println("Is false")
	}

	if value := true; value {
		fmt.Println("Is true")
	}

	number := 3

	if number == 1 {
		fmt.Println("One")
	} else if number == 2 {
		fmt.Println("Two")
	} else if number == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("Other")
	}

	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Undefined number")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Undefined number")
	}

	switch {
	case number > 0:
		fmt.Println("Positive")
	case number < 0:
		fmt.Println("Negative")
	case number == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Undefined number")
	}
}
