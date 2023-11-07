package router

import (
	"Cargo/Cargo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/deneme", controller.CargosWorkmans)

	return r

}

func deneme(c *gin.Context) {
	c.JSON(200, gin.H{
		"deneme": "deneme",
	})
}
