package service

import (
	"example/promote/model"
	"example/promote/repository"
)

func GetPromotion(id int) model.Promotion {
	promotion := repository.GetById(id)
	return promotion
}

func GetPromotions(offset any, limit any) []*model.Promotion {
	promotions := repository.GetAll(offset, limit)
	return promotions
}

func AddPromotion(promotion model.Promotion) {
	repository.Add(promotion)
}

func AddPromotions(promotions []*model.Promotion) {
	repository.AddAll(promotions)
}

func DropTable() {
	repository.DropTable()
}
