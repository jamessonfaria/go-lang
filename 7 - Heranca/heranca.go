package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Bola", "Fran", 22, 177}
	fmt.Println(p1)

	e1 := estudante{p1, "Curso A", "Univida"}
	fmt.Println(e1)

	e2 := estudante{pessoa{"Caique", "Silva", 40, 189}, "Curso A", "Univida"}
	fmt.Println(e2.nome)

}