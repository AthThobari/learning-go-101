package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   // Standard field for the primary key
	Name      string // A regular string field  // Uses sql.NullTime for nullable time fields
	Age       int
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
type Student struct {
	Name      string 
	Age       int
	Class     string
	CreatedAt time.Time
	UpdatedAt time.Time 
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/mengenal_sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("There is something error in database connection", err)
	}

	fmt.Println("Database connected")

	//Melakukan pembuatan model
	db.AutoMigrate(&User{})

	// data_1 := []User{
	// 	{Name: "Afif",
	// 	Age: 16,},
	// 	{Name: "Joko",
	// 	Age: 16,},
	// 	{Name: "Eko",
	// 	Age: 16,},
	// 	{Name: "Bambang",
	// 	Age: 16,},
	// }

	// db.Create(&data_1)

	// data := []User{}

	// db.Find(&data)

	// db.Where("name = ?", "Bambang").Find(&data)

	// for _, el := range data {

	// 	fmt.Println("Nama:", el.Name)
	// 	fmt.Println("Umur:", el.Age)
	// }


	//Membuat tabel dan data student
	db.AutoMigrate(&Student{})

	data_student := []Student{}

	// db.Create(data_student)

	fmt.Println("=== Mencetak data student dengan umur diatas 10 ===")

	db.Raw("SELECT name, age, class FROM students WHERE age > ?", 10).Scan(&data_student)
	for _, el := range data_student {
		fmt.Println("Nama:", el.Name)
		fmt.Println("Umur:", el.Age)
		fmt.Println("Kelas:", el.Class)}


   	fmt.Println("=== Mencetak data student di class tertentu ===")

	db.Where("class = ?", "RPL").Find(&data_student)
	for _, el := range data_student {
		fmt.Println("Nama:", el.Name)
		fmt.Println("Umur:", el.Age)
		fmt.Println("Kelas:", el.Class)
	}



}
