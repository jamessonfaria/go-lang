package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("devtools@gmail.com")
	fmt.Println(erro)
}

// 1- go mod init nome_modulo  	---------> Inicia o modulo e gera o arquivo go.mod
// 2- go build				 	---------> Cria o arquivo executavel do codigo