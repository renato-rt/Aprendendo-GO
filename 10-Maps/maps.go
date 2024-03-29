package main

import "fmt"

func main() {
	fmt.Println("Maps")
//Map só aceita que seja um tipo todos os campos
	usuario := map[string]string {
		"nome":			"Pedro",
		"sobrenome":	"Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome" : {
			"primeiro": "João",
			"ultimo" :	"Pedro",
		},
		"curso" : {
			"nome" : "Egenharia",
			"Prédio": "1",
		},
	}

	fmt.Println(usuario2)
	delete (usuario2, "nome")
	fmt.Println(usuario2)
	
	usuario2["signo"] = map[string]string{
		"nome":"Gemeos",
	}
	fmt.Println(usuario2)
}