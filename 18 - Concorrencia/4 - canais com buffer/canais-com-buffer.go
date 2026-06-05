package main

import "fmt"

func main(){

	// quando vc cria um canal na mesma funcao dessa forma vc toma erro deadlock!
	// canal := make(chan string)
	// o proximo parametro do make é o buffer, quando vc informa o buffer vc consegue chamar o chanel na mesma funcao
	// como tem 2 buffers é possivel mandar duas mensagens para o canal e receber as mensagens 
	canal := make(chan string, 2) 
	canal <- "Ola mundo"
	canal <- "Programando em go"

	mensagem := <- canal
	mensagem2 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}