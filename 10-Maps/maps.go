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

	usuario2 := map[string]map[string]string(
		"nome:": {
				"primeiro": "João"
		}
	)
}