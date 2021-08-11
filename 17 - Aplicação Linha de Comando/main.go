package main

import (
	"linha-de-comando/App"
	"log"
	"os"
)

func main() {
	aplicacao := App.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
