package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")

	retorno := func (texto string) string{
		fmt.Println(texto)
		//PAra passar o parametro de forma abreviada, basta colocar %s para strings e %d para numeros
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando Parametro")

	fmt.Println(retorno)
}