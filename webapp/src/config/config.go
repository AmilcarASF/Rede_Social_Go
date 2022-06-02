package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL é a URL de comunicação com a Api
	APIURL = ""
	// Porta da aplicação web
	Porta = 0
	// HashKey é utilizado para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizado para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as váriaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
