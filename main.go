package main

import (
	//  "errors"
	"court/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	servPort := GetArg("PORT", "8080")
	db := OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'dbo'")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var res string
	for rows.Next() {
		rows.Scan(&res)
		fmt.Println(res)
	}

	http.HandleFunc("/", handlers.Redirect)
	err = http.ListenAndServe(":"+servPort, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
