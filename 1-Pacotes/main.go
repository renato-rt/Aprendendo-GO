package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

//Escrever registra uma mensagem na tela
func main() {
	fmt.Println("Escrevendo do Arquivo main")
	auxiliar.Escrever()

	//go mod tidy para remover dependencias não usadas

	erro := checkmail.ValidateFormat("tomikawa.renatomail.com")
	fmt.Println(erro)
}
