package main

import "fmt"

func calculosMatematicos(num1, num2 int) (soma int, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

func main(){
	soma, sub := calculosMatematicos(20,30)
	fmt.Println(soma, sub)
}