package main

import (
	"fmt"
	"time"
)

func main() {

	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j", j, "sapo")
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indeice, nome := range nomes {
		fmt.Println(indeice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	type usuarioStruct struct {
		nome      string
		sobrenome string
	}

}
