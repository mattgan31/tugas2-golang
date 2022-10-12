package routers

import (
	"golang-tugas2/controllers"
	"golang-tugas2/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	db := database.GetDB()
	router := gin.Default()
	ctr := controllers.New(db)
	router.GET("/orders", ctr.ShowOrders)
	router.POST("/orders", ctr.CreateOrder)
	router.DELETE("orders/:orderID", ctr.DeleteOrder)
	router.PUT("orders/:orderID", ctr.UpdateOrder)

	return router
}
