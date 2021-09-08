package main

import (
	"fmt"
	"time"
)

func main() {
	//Para realizar uma goroutine apenas coloque a palavra "go" na frente da chamada
	go escrever("Olá Mundo!")
	escrever("Programando em GO")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time. Sleep(time.Second)
	}
}
