package main

import "fmt"

// generic recebe uma interface vazia, mas atualmente o golang permite o tipo any que pode receber qualquer coisa
//func generica(interf interface{})  {
func generica(interf any)  {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Generics")

	generica("String")
	generica(1)
	generica(true)

	// essa funcao Println recebi a interface{} que hoje em dia pode ser any
	fmt.Println("String", true, 123)

	mapa := map[any]any{
		1: "String",
		true: float32(100),
		"String": "String",
	}

	fmt.Println(mapa)

}