package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func connectDB() *sql.DB {

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", "localhost", "SA", "J55fonteles", "1433")
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Erro ao conectar com o BD:", err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Erro ao conectar com o BD:", err)
	}

	return db


}