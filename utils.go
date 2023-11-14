package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetArg(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
func OpenDB() *sql.DB {
	user := GetArg("PG_USER", "postgres")
	passwd := GetArg("PG_PASSWORD", "postgres")
	port := GetArg("PG_PORT", "5432")
	host := GetArg("PG_HOST", "localhost")
	DBName := GetArg("PG_DBNAME", "court")

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
