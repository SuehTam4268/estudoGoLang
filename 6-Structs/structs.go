package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "davi"
	u.idade = 21
	fmt.Println(u)
	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Math", 17, enderecoExemplo}
	fmt.Println(usuario2)
}
