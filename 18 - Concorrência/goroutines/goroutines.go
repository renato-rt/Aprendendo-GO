package main

import (
	"fmt"
	"time"
)

func main() {
	escrever("Olá Mundo!")
	escrever("Programando em GO")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time. Sleep(time.Second)
	}
}
