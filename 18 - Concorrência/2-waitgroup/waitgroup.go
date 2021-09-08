package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)
//Rotina 1
	go func () {
		escrever("Hello")
		waitgroup.Done()// Retira a rotina do grupo após executar.
	}()
//Rotina 2
	go func() {
		escrever("in")
		waitgroup.Done()
	}()
//Rotina 3
	go func() {
		escrever("Good")
		waitgroup.Done()
	}()
//Rotina 4
	go func() {
		escrever("World")
		waitgroup.Done()
	}()

	waitgroup.Wait() //Manda esperar a contagem de rotinas chegar em zero.

}

func escrever(texto string) {
	for i := 0; i < 5 ; i++ {
		fmt.Println(texto)
		time. Sleep(time.Second)
	}
}