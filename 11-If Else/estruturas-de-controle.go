package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
	fmt.Println("Maior")
	}else {
		fmt.Println("Menor ou igual")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}else if outroNumero < -10 {
		fmt.Println("Numero é menor que -10")
	}else {
		fmt.Println("Entre 0 e -10")
	}


}