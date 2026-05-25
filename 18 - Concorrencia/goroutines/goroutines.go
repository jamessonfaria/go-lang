package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	// PARALELISMO: SAO TAREFAS EXECUTADAS AO MESMO TEMPO E SO PODE OCORRER EM MULTI CORE (PROCESSADORES)
	// CONCORRENCIA: SAO TAREFAS QUE NAO NECESSARIAMENTE ESTAO EXECUTANDO AO MESMO TEMPO, EM UM MONO CORE (PROCESSADOR)
	//     A TAREFA 1 EXECUTARIA UMA PARTE, DEPOIS A TAREFA 2 EXECUTARIA OUTRA PARTE E ASSIM ATE FINALIZAR

	// Concorrência = várias tarefas “avançando juntas”, mas não necessariamente ao mesmo tempo.
	// 		Na programação:
	//			um processador pode alternar entre tarefas muito rápido;
	// 			isso dá a impressão de simultaneidade.
	// Paralelismo = várias tarefas executando literalmente ao mesmo tempo.
	//		Na programação:
	// 			múltiplos núcleos da CPU executam tarefas ao mesmo tempo.

	go escrever("Ola mundo") // goroutine, ele nao espera o programa terminar e segue para a proxima linha
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
