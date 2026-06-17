package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o MySQL
)

func Conectar() (*sql.DB, error) {
	stringConexao := "my_user:my_user_password@/my_database?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	fmt.Println("Conexão esta aberta!")
	return db, nil
}