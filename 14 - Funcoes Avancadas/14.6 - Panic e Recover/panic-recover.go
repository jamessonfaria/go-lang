package main

import "fmt"

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperaExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXTAMENTE 6 !!!") // depois do panic e chamado o defer e ai a funcao recover
}

func main() {
	fmt.Println(alunoEstaAprovado(5, 8))
	fmt.Println("Pós execução 1!")

	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução 2!")

}	