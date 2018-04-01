package main

import (
	"github.com/gin-gonic/gin"
	"github.com/avegao/iot-fronius-push-service/controller"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	initFroniusRouter(router)

	return router
}


func initFroniusRouter(router *gin.Engine) {
	froniusRouter := router.Group("/fronius")

	initFroniusCurrentDataRouter(froniusRouter)
}

func initFroniusCurrentDataRouter(router *gin.RouterGroup) {
	currentDataRouter := router.Group("/current_data")
	currentDataRouter.POST("/meter", froniusController.PostCurrentDataMeterAction)
	currentDataRouter.POST("/inverter", froniusController.PostCurrentDataInverterAction)
	currentDataRouter.POST("/powerflow", froniusController.PostCurrentDataPowerflowAction)
	currentDataRouter.POST("/sensor_card", froniusController.PostCurrentDataSensorCardAction)
	currentDataRouter.POST("/string_control", froniusController.PostCurrentDataStringControlAction)
}
