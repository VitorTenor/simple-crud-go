package db

import (
	"database/sql"
	"fmt"
	"simpleCrudGo/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	dbConfig := configs.GetDBConfig()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
