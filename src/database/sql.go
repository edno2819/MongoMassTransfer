package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SqlConection() {
	// Define a string de conexão (substitua pelos seus dados)
	dsn := "usuario:senha@tcp(127.0.0.1:3306)/nome_do_banco"

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	defer db.Close()

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	fmt.Println("Conectado ao MySQL com sucesso!")

	// Exemplo: realizando uma consulta
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal("Erro ao executar a consulta:", err)
	}

	fmt.Println("Versão do MySQL:", version)
}
