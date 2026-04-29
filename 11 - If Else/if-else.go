package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")
	numero := 30
	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("É menor ou igual que 15")
	}

	// IF INIT, interessante pois a variavel outroNumero é limitada a esse escopo do if apenas, depois ela deixa de existir
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("É maior que zero.")
	} else if outroNumero == 0 {
		fmt.Println("É zero.")
	} else {
		fmt.Println("É menor ou igual a zero.")
	}
}