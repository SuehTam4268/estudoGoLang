package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}
func funcao2() {
	fmt.Println("Executando a função 2")
}
func alunosEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada.")
	fmt.Println("Entrando na função de verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {

		return true
	}

	return false
}

func main() {

	fmt.Println(alunosEstaAprovado(7, 8))
}
