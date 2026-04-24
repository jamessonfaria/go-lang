package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1 // atribuicao de um valor é uma copia, enderecos de memoria diferentes

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA
	var var3 int = 100
	var ponteiro *int // com o * na frente do tipo se declara um ponteiro
	fmt.Println(var3, ponteiro)
	
	ponteiro = &var3 // com o & passa o endereco de memoria de uma var para o ponteiro
	fmt.Println(var3, ponteiro)

	fmt.Println(var3, *ponteiro) // colocando o * na frente do ponteiro vc ver o valor, desreferenciacao

	var3 = 2000
	fmt.Println(var3, *ponteiro)
}