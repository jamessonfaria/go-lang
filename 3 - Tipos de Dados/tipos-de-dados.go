package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000000000
	var numero2 int = 10000000000000000

	fmt.Println(numero)
	fmt.Println(numero2)

	// uint - unsigned int, int sem sinal
	var numero3 uint = 1000
	fmt.Println(numero3)

	// alias
	// INT32 = rune
	// INT8 = byte
	var numero4 rune = 123123123
	var numero5 byte = 10
	
	fmt.Println(numero4)
	fmt.Println(numero5)

	var num6 float32 = 12222.3222
	fmt.Println(num6)

	num7 := 22222.33
	fmt.Println(num7)

	var str = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// nao tem o tipo char, ele faz por inferencia
	char := 'B'
	fmt.Println(char)

	// todo tipo tem um valor padrao, "" - string, 0 - int, false - boolean, <nil> - error
	var texto3 string
	fmt.Println(" >>>", texto3)

	var boolean1 bool
	fmt.Println(boolean1)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Error interno")
	fmt.Println(erro2)

}