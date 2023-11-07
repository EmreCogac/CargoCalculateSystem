package controller

import (
	"Cargo/Cargo/database"
	"Cargo/Cargo/initializers"
	"Cargo/Cargo/models"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	config, err := initializers.LoadEnv("..")
	if err != nil {
		log.Fatal("hata var controller ")
	}
	database.ConnectDB(&config)

}
func GetAll(c *gin.Context) {
	db := database.GlobalDB
	Cok := []models.Cok{}

	db.Preload("Tek").Find(&Cok)

	c.JSON(200, gin.H{
		"a": Cok,
	})

}

func CargosWorkmans(c *gin.Context) {
	db := database.GlobalDB
	cargo := []models.Cargos{}
	workmans := []models.Workmans{}

	db.Preload("WorkCargo").Find(&cargo).Find(&workmans)

	c.JSON(200, gin.H{
		"a": workmans,
		"b": cargo,
	})

}

// func Cargo(ctx *gin.Context) {

// 	db := database.GlobalDB
	
// }
