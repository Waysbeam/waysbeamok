package database

import (
	"Be_waysbeam/models"
	"Be_waysbeam/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Transaction{},
		&models.Profile{},
	)
	if err!= nil{
		fmt.Println(err)
		panic("Migration Failed")
	}
	fmt.Println(("Migration Success"))
}