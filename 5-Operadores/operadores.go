package main

import "fmt"

func main() {

	//Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Fim Aritméticos

	//Atribuição
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")
	var variavel1 string = "string"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//Fim Atribuição

	//Operadores Relacionais
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//Fim Relacionais

	//Lógicos
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//Fim Lógicos

	//Unários
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //Numero = numero - 20
	fmt.Println(numero)

	numero *= 3 //numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	//Fim Unários

	//Ternários
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
