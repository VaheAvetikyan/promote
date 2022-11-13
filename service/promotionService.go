package service

import (
	"example/promote/model"
	"example/promote/repository"
)

type PromotionService interface {
	GetPromotion(id int) model.Promotion
	GetPromotions(offset string, limit string) []*model.Promotion
	AddPromotion(promotion model.Promotion)
	AddPromotions(promotions []*model.Promotion)
	DropTable()
}

type service struct{}

func NewPromotionService() PromotionService {
	return &service{}
}

func (*service) GetPromotion(id int) model.Promotion {
	promotion := repository.GetById(id)
	return promotion
}

func (*service) GetPromotions(offset string, limit string) []*model.Promotion {
	promotions := repository.GetAll(offset, limit)
	return promotions
}

func (*service) AddPromotion(promotion model.Promotion) {
	repository.Add(promotion)
}

func (*service) AddPromotions(promotions []*model.Promotion) {
	repository.AddAll(promotions)
}

func (*service) DropTable() {
	repository.DropTable()
}
