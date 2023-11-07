package main

import (
	"Cargo/Cargo/database"
	"Cargo/Cargo/initializers"
	"Cargo/Cargo/models"
	"fmt"
	"log"
)

func init() {

	config, err := initializers.LoadEnv("../..")
	if err != nil {
		log.Fatal("fatal error path from migrate", err)
	}
	database.ConnectDB(&config)
}

func main() {
	database.GlobalDB.AutoMigrate(models.Cok{})
	database.GlobalDB.AutoMigrate(models.Tek{})
	database.GlobalDB.AutoMigrate(models.Workmans{})
	database.GlobalDB.AutoMigrate(models.Vehicles{})
	database.GlobalDB.AutoMigrate(models.Cargos{})
	database.GlobalDB.AutoMigrate(models.WorkCargo{})
	database.GlobalDB.AutoMigrate(models.Packages{})
	database.GlobalDB.AutoMigrate(models.Admins{})
	fmt.Println("migration is compeleted")

}
