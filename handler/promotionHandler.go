package handler

import (
	"encoding/csv"
	"example/promote/model"
	"example/promote/service"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecords(context *gin.Context) {
	offset, exists := context.GetQuery("offset")
	if !exists {
		offset = "0"
	}
	limit, exists := context.GetQuery("limit")
	if !exists {
		limit = "100"
	}
	context.IndentedJSON(http.StatusOK, service.GetPromotions(offset, limit))
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

func UploadPromotions(context *gin.Context) {
	fileHeader, err := context.FormFile("file")
	if err != nil {
		log.Print("Error getting the fileHeader. ", err)
		return
	}
	log.Printf("Uploaded file Header name: %+v\n", fileHeader.Filename)
	log.Printf("Uploaded file Header size: %+v bytes\n", fileHeader.Size)
	log.Printf("File mime type: %+v\n", fileHeader.Header)

	service.DropTable()

	file, err := fileHeader.Open()
	defer file.Close()
	reader := csv.NewReader(file)
	rows := make([]*model.Promotion, 0)
	for i := 0; ; i++ {
		csvLine, err := reader.Read()
		if i%100 == 0 && i != 0 {
			service.AddPromotions(rows)
			rows = make([]*model.Promotion, 0)
		}
		// Stop at EOF.
		if err == io.EOF {
			service.AddPromotions(rows)
			rows = nil
			break
		}
		if err != nil {
			log.Panic(err)
		}
		price, _ := strconv.ParseFloat(csvLine[1], 32)
		rows = append(rows, &model.Promotion{
			Id:             csvLine[0],
			Price:          float32(price),
			ExpirationDate: csvLine[2],
		})
	}
	context.IndentedJSON(http.StatusOK, "Uploaded Successfully.")
}
