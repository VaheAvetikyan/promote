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
	router.GET("/promotions/:id", handler.GetPromotionById)
	router.POST("/promotions", handler.AddPromotion)
	return router
}

func main() {
	config, err := configuration.New()
	if err != nil {
		log.Panic("Configuration Error", err)
	}
	router := Routes()
	err = router.Run(config.HOST + ":" + config.PORT)
	if err != nil {
		return
	}
}
