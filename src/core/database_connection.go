package core

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

func InitDB() (*sql.DB, error) {
	dsn := "alexia:alexia2005(localhost:3306)/mate_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al verificar conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
	return db, nil
}
