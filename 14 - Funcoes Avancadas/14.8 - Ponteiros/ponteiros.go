package main

import "fmt"

func inverteSinal(num int) int {
	return num * -1
}

func inverteSinalComPonteiro(numAddress *int) {
	*numAddress = *numAddress * -1
}

func main()  {
	fmt.Println("Ponteiros")
	variavel := 20
	retornoVar := inverteSinal(variavel)
	fmt.Println(variavel)
	fmt.Println(retornoVar)

	novaVariavel := 40
	fmt.Println(novaVariavel)
	inverteSinalComPonteiro(&novaVariavel)
	fmt.Println(novaVariavel)

}