package main

import (
	"example/promote/configuration"
	"example/promote/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/promotions", handler.GetPromotion)
	router.GET("/promotions/:id", handler.GetPromotionById)
	router.POST("/promotions", handler.AddPromotion)

	//curl -X POST http://localhost:1321/promotions/upload -F "file=@/<filepath>" -H "Content-Type: multipart/form-data"
	router.POST("/promotions/upload", handler.UploadPromotions)
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
		log.Fatal("Could not initialize Server.")
	}
}
