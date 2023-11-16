package main

import (
	"court/handlers"
	"court/model"
	"court/utils"
	"fmt"
	"net/http"
	"os"
)

func main() {

	servPort := utils.GetArg("PORT", "8080")
	db := model.OpenDB()
	defer db.Close()
	model.SetupTable(db)
	http.HandleFunc("/", handlers.Redirect)
	err := http.ListenAndServe(":"+servPort, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
