package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2) // estou informando que terei dois waits

	go func(){ // goroutine
		escrever("Ola mundo")
		waitgroup.Done() // quando terminar de rodar a funcao acima ele diminui -1 do contador do wait group
	}()

	go func(){ // goroutine
		escrever("Programando em Go!")
		waitgroup.Done()  // quando terminar de rodar a funcao acima ele diminui -1 do contador do wait group
	}()

	waitgroup.Wait() // ele espera ate o grupo das goroutines chegar em 0, espera as duas terminarem
}

func escrever(texto string) {
	for i:=0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
