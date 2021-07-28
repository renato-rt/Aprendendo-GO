package main

import (
	"errors"
	"fmt"
)

func main() {
	//com sinal tipos int, int8. int16, int32, int64
	numero := 1000000000000
	fmt.Println(numero)
	//sem sinal
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//rune é igual a int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte é igual a uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//Fim Numeros Reais

	//Strings
	var str string = "Teste"
	fmt.Println(str)
	str2 := "Teste2"
	fmt.Println(str2)

	//Tabela ASK
	char := 'B'
	fmt.Println(char)
	//Fim Strings

	//String vazia tem valor de zero
	var texto string
	fmt.Println(texto)

	//Booleano
	var booleanol bool = false
	fmt.Println(booleanol)

	//Erro
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
