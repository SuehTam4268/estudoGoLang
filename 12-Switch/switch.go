package main

import (
	"fmt"
)

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "segunda Feira"
	case 3:
		return "terça feira"
	case 4:
		return "quarta feira"
	case 5:
		return "Quinta feira"
	case 6:
		return "Sexta feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"

	}
}

func diaDaSemana2(numero int)string{
	switch{
	case numero == 1:
		return "Domingo"
	case numero == 12:
		return "segunda Feira"
	case numero == 13:
		return "terça feira"
	case numero == 14:
		return "quarta feira"
	case numero == 1 5:
		return "Quinta feira"
	case numero == 16:
		return "Sexta feira"
	case numero == 17:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(8)
	fmt.Println(dia)
}
