package main

import (
	"fmt"
	"mengenal_sql/config"
)

func main() {
	fmt.Println("Connecting to database")
	db := config.NewDatabase()

	if db == nil {
		fmt.Println("Failed connecting to database")
		return
	}

	fmt.Println("Main function has been trigerred")
}
