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

	// db.Where("id = ?", 1).Updates(&config.Student{
	// 	Name: "joko",
	// 	Age: 30,
	// })

	//Query update name 'Joe' diganti dengan nama 'Hans'
	db.Where("name = ?", "Joe").Updates(&config.Student{
		Name: "Hans",
	})
}
