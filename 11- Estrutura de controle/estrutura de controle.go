package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor que zero")
	}
}
