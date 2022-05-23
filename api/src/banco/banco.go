package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		fmt.Println("Erro sql.Open")
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro db.Ping")
		db.Close()
		return nil, erro
	}

	return db, nil
}
