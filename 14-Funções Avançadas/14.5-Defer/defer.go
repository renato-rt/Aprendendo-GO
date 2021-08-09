package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}
func alunoEstaAprovado(n1, n2 float32) bool {
	
	//Defer = adiar
	defer fmt.Println("Média calculada, Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + (n2 * 2)) / 3

	if media >= 5 {
		fmt.Println("Média calculada, resultado será retornado, Aluno Aprovado")
		return true
	}
	fmt.Println("Média calculada, resultado será retornado, Aluno Reprovado")
	return false
}

func main() {
	funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 10))
}