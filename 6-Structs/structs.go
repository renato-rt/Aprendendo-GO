package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	//Primeira forma de declarar arquivo Structs
	var u usuario
	u.nome = "Renato"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua 1", 0}

	//Fim

	//Segunda forma
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)
	//Fim

	//Terceira forma
	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
	//Fim
}
