package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Entrando na funcao para verificar se o aluno esta aprovado ou nao")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main(){
	// DEFER = adiar ate a ultima execucao possivel dentro do bloco
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(10, 2))
}
