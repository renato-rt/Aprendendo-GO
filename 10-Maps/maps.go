package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome":			"Pedro",
		"sobrenome":	"Silva",
	}
	fmt.Println(usuario)
}