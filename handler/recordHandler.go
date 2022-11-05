package handler

import (
	"example/promote/service"
	"net/http"
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
		return
	}
	context.IndentedJSON(http.StatusOK, service.GetRecord(id))
}
