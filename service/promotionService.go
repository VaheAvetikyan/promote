package service

import (
	"example/promote/model"
	"example/promote/repository"
)

func GetPromotion(id int) model.Promotion {
	promotion := repository.GetById(id)
	return promotion
}

func GetPromotions() []*model.Promotion {
	promotions := repository.GetAll()
	return promotions
}

func AddPromotion(promotion model.Promotion) {
	repository.Add(promotion)
}
