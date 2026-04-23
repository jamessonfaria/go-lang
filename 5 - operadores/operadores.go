package main

import "fmt"

func main() {

	// =======>  + - / * %

	soma := 1 + 2
	sub := 2 - 3
	mult := 10 * 5
	div := 10 / 4
	resto := 10 % 2

	fmt.Println(soma, sub, mult, div, resto)

	var num1 int16 = 10
	var num2 int16 = 25

	soma2 := num1 + num2
	fmt.Println(soma2)

	 // FIM DOS ARITIMETICOS

	var var1 string = "string"
	var2 := "string2"
	fmt.Println(var1, var2)

	// FIM DOS OPERADORES ATRIB

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// FIM DOS RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println("------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero *= 20
	fmt.Println(numero)
	numero /= 100
	fmt.Println(numero)
	// FIM OPERADORES UNARIOS
	
	// OPERADOR TERNARIO NAO EXISTE NO GO, TEM QUE USAR IF ELSE NORMAL


}
