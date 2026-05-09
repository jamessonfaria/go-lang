package main

import "fmt"

func closure() func() { // funcao que retorna uma funcao
	texto := "Dentro da funcao closure"

	funcao := func() {		// funcao que sera retornada
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}