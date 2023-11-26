package main

import (
	"fmt"
	"log"
	"os"
	"search-ip/app"
)

func main() {
	fmt.Println("Run app")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
