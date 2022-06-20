package main

import "fmt"

func main() {
	var variavel1 string = "bananas de pijama"
	variavel := 1 + 1
	fmt.Println(variavel1)
	fmt.Println(variavel)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	A, M := "Amanda", "Matheus"
	fmt.Println(A, M)
	A, M = M, A
	fmt.Println(A, M)

}
