package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	ativo bool
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs.")

	var usu1 usuario 
	fmt.Println(usu1)

	usu1.nome = "James"
	usu1.idade = 20
	usu1.ativo = true
	fmt.Println(usu1)

	end := endereco{"rua x", 1}
	usu2 := usuario{"Carlos", 12, false, end}
	fmt.Println(usu2)

	usu3 := usuario{nome: "Fredinho"}
	fmt.Println(usu3)


}