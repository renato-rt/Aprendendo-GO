package main

import (
	"fmt"
	"time"
)

func Fibonacci(posicao uint) uint{
	if posicao <= 1 {
		return posicao
	}

	return Fibonacci(posicao -2)+ Fibonacci(posicao -1)
}
func main() {
	fmt.Println("Funções Rescursivas")

	///Função que vai retornar o numero que está numa determinada posição da sequencia de Fibonacci, é uma sequencia de numeros onde o proximo numero é sempre a soma dos dois numeros anteriores.

	posicao := uint(100)

	for i := uint(1); i < posicao; i++ {
		time.Sleep(time.Second)
		fmt.Println(Fibonacci(i))
	}
	fmt.Println((Fibonacci(posicao)))
}