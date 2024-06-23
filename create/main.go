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

	//create new data

	//Create data student
	student_1 := []config.Student{
		{
			Name: "John",
			Age:  25,
		},
		{
			Name: "Thomas",
			Age: 30,
		},
		{
			Name: "Jane",
			Age: 18,
		},
		{
			Name: "Joe",
			Age: 17,
		},
	}

	db.Create(&student_1)

	fmt.Println("Data has been created")
}
