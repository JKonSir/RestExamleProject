package support

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func (pgDb *PostgresDB) InitDB(user, password, dbname string)  {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			user, password, dbname)

	var err error
		pgDb.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}