package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Carlos",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["curso"])
	fmt.Println(usuario2["curso"]["campus"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"id": "1",
		"nome": "Peixe",
	}
	fmt.Println(usuario2)

}