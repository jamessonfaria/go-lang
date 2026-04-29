package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero  {
	case 1:
		return "Domingo"
	case 2: 
		return "Segunda-feira"
	case 3: 
		return "Terça-Feira"
	default: 
		return "Numero Invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // faz ele retorar o bloco do proximo case sem avaliar a condicao
	case numero == 2: 
		diaDaSemana = "Segunda-feira"
	case numero == 3: 
		diaDaSemana = "Terça-Feira"
	default: 
		diaDaSemana = "Numero Invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	
	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}