package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/utils"
)

//func init() {
//	hashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//	fmt.Println(hashkey)

//	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//	fmt.Println(blockkey)
//}

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Rodando WebApp na porta %d...\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
