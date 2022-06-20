package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000000000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	var numero3 rune = 12456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	var STR string = "texto"
	fmt.Println(STR)

	STR2 := "Texto 2"
	fmt.Println(STR2)

	char := 'A'
	fmt.Println(char)

	var boleano1 bool
	fmt.Println(boleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
