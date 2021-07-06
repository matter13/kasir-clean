<<<<<<< HEAD
package config

import (
	//"go-clean/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dbs := "root:@tcp(127.0.0.1:3306)/db_kasir?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dbs), &gorm.Config{})

	if err != nil {
		panic("Database tidak ada")
	} else {
		fmt.Println("Terkoneksi")
	}
	return db
}
=======
package config

import (
	//"go-clean/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dbs := "root:@tcp(127.0.0.1:3306)/db_kasir?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dbs), &gorm.Config{})

	if err != nil {
		panic("Database tidak ada")
	} else {
		fmt.Println("Terkoneksi")
	}
	return db
}
>>>>>>> master
