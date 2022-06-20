package enderecos

import "strings"

//tipo endereco
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	//enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemumTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemumTipoValido = true
		}
	}
	if enderecoTemumTipoValido {
		return primeiraPalavraDoEndereco
	}
	return "Tipo Invalido"
}
