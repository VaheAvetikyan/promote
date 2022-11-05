package main

import (
	"example/promote/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/records", handler.GetRecords)
	router.GET("/records/:id", handler.GetRecordById)
	router.Run("localhost:8080")
}
