package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2JSON := `{"nome": "Toby", "raca": "Podle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}	

	fmt.Println(c2)
}