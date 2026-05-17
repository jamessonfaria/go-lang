package main

import "fmt"

var n int 

// funcao init é sempre executada primeiro em cada arquivo
func init()  {
	fmt.Println("Rodando funcao init")
	n = 10
}

func main()  {
	fmt.Println("Rodando o main")
	fmt.Println(n)

	
}