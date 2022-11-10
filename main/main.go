package main

import (
	"example/promote/configuration"
	"example/promote/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/promotions", handler.GetRecords)
	router.GET("/promotions/:id", handler.GetRecordById)
	router.POST("/promotions", handler.AddRecord)
	return router
}

func main() {
	configuration, err := configuration.New()
	if err != nil {
		log.Panic("Configuration Error", err)
	}
	router := Routes()
	router.Run(configuration.HOST + ":" + configuration.PORT)
}
