package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (user usuario) salvar() {
	fmt.Printf("Salvando o usuario: %s no banco de dados\n", user.nome)
}

func (user usuario) maiorDeIdade() bool{
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade++
}

func main()  {
	fmt.Println("Metodos")
	usuario1 := usuario{"Jamesson", 46}
	fmt.Println(usuario1)
	usuario1.salvar()

	maiorIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)


}