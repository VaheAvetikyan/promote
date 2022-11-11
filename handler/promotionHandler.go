package handler

import (
	"example/promote/model"
	"example/promote/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecords(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, service.GetPromotions())
}

func GetPromotionById(context *gin.Context) {
	identifier := context.Param("id")
	id, err := strconv.Atoi(identifier)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provided id is invalid.",
		})
		log.Fatal("Provided id is invalid. ", err)
	}
	promotion := service.GetPromotion(id)
	if promotion == (model.Promotion{}) {
		context.IndentedJSON(http.StatusNotFound, fmt.Sprintf("Record not found with id: %s", identifier))
	} else {
		context.IndentedJSON(http.StatusOK, promotion)
	}
}

func AddPromotion(context *gin.Context) {
	promotion := model.Promotion{}
	err := context.Bind(&promotion)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provided data is not valid.",
		})
		log.Fatal("Provided data is not valid. ", err)
	}
	service.AddPromotion(promotion)
	context.IndentedJSON(http.StatusOK, "Added Successfully.")
}
