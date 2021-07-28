package main

import "fmt"

func main() {
	var variavel1 string = "variável 1975"
	//Variaveis ou bibliotecas nao utilizadas dão erro no programa!
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "45678"
		variavel4 string = "1234"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "\nVariável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
