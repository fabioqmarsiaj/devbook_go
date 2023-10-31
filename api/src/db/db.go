package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Connect open connection with DB and return it
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
