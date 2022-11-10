package handler

import (
	"example/promote/service"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecords(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, service.GetRecords())
}

func GetRecordById(context *gin.Context) {
	identifier := context.Param("id")
	id, err := strconv.Atoi(identifier)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provided id is invalid",
		})
		return
	}
	context.IndentedJSON(http.StatusOK, service.GetRecord(id))
}

func AddRecord(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	context.IndentedJSON(http.StatusOK, extension)
}
