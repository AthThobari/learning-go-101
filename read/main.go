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

	var student []config.Student

	//Query cari data dengan 'Age' < 19
	fmt.Println("======Query cari data dengan 'Age' < 19======")
	
	db.Where("age < ?", 19).Find(&student)
	for _, data := range student {
		fmt.Println("ID:", data.ID)
		fmt.Println("Nama:", data.Name)
		fmt.Println("Umur:", data.Age)
	}

	// Query cari data dengan kata nama yang mengandung kata "jo"
	fmt.Println("======Query cari data dengan kata nama yang mengandung kata \"jo\"======")

	searchQuery := "%" + "jo" + "%"
	db.Where("name LIKE ?", searchQuery).Find(&student)
	for _, data := range student {
		fmt.Println("ID:", data.ID)
		fmt.Println("Nama:", data.Name)
		fmt.Println("Umur:", data.Age)
	}



	// searchQuery := "%" + "Bu" + "%"

	// db.Limit(3).Where("name LIKE ?", searchQuery).Find(&student)

	// for _, data := range student {
	// 	fmt.Println("ID:", data.ID)
	// 	fmt.Println("Nama:", data.Name)
	// 	fmt.Println("Umur:", data.Age)
	// }
}
