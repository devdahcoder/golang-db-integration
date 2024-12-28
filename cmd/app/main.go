package main

import (
	"fmt"
	"golang-db-integration/internals/database"
)

func main() {
	db := database.GetPsqlDatabase()
	defer func() {
		if err := db.ClosePsqlDb(); err != nil {
			fmt.Println(fmt.Sprintf("Error closing database: %v", err))
		}
	}()
}
