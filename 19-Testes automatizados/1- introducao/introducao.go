package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	tipoEndeco := enderecos.TipoDeEndereco("ADU")
	fmt.Println(tipoEndeco)
}
