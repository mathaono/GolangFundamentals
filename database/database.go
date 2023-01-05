package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o banco
)

func Connect() (*sql.DB, error) {
	const PostgresDriver = "postgres"

	const User = "postgres"

	const Host = "localhost"

	const Port = "5432"

	const Password = "adv17667"

	const DbName = "devbook"

	const TableName = "usuarios"

	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	//urlConnect := "host=localhost port=5432 user=postgres password=adv17667 dbname=devbook sslmode=disable"

	db, err := sql.Open("postgres", DataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
