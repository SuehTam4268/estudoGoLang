package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	fmt.Println(numeros)
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	fmt.Println("Funções variaticas")
	totalDaSoma := soma(9, 10, 11)
	fmt.Println(totalDaSoma)

	escrever("Olá, mundo", 10, 2, 3, 4, 5, 6, 7)
}
