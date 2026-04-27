package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAY É RIDIGO, POUCO USADO NO GO 
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	fmt.Println(array1)

	array2 := [4]string{"Pos1", "pos2", "Pos3", "pos4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE NÃO É UM ARRAY
	slice := []int{1, 2, 3}
	slice = append(slice, 1)
	slice = append(slice, 2000)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array3[1:3]
	fmt.Println(slice2)
	slice2 = append(slice2, 1000)
	fmt.Println(slice2)

	// ponteiro, o valor do array mudou e como era um ponteiro mudou para o slice tambem
	array3[1] = 50
	fmt.Println(slice2)


	// ARRAYS INTERNOS
	fmt.Println("---------------------------------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	// len = 10 → quantos elementos você pode usar agora
	// cap = 11 → quanto espaço total existe antes de realocar
	
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) //capacidade

}