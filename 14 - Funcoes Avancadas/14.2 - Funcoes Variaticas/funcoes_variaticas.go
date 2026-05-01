package main

import "fmt"

func totalizandoNumeros(texto string, numeros ...int) {
	total := 0
	for _, numero := range (numeros){
		total += numero
	}
	fmt.Println(texto, total)
}

func main()  {
	totalizandoNumeros("Total é : ", 1,2,3,4,5,6)
}