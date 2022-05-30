package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando WebApp na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
