package model

import (
	"court/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenDB() *sql.DB {
	user := utils.GetArg("PG_USER", "postgres")
	passwd := utils.GetArg("PG_PASSWORD", "postgres")
	port := utils.GetArg("PG_PORT", "5432")
	host := utils.GetArg("PG_HOST", "localhost")
	DBName := utils.GetArg("PG_DBNAME", "court")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, passwd, DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SetupTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS court (id SERIAL PRIMARY KEY, long_link TEXT NOT NULL, short_link TEXT NOT NULL );")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO court (long_link, short_link) VALUES ('https://google.com/', '/000000')")
	if err != nil {
		log.Fatal(err)
	}
}
