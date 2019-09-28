package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() *sql.DB {

	dsn := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%s", "localhost", "AdventureWorks2017", "SA", "J55fonteles", "1433")
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Erro ao conectar com o BD:", err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Erro ao conectar com o BD:", err)
	}

	return db


}