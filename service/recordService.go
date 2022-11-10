package service

import (
	"example/promote/model"
	"example/promote/repository"
)

func GetRecord(id int) model.Promotion {
	promotion := repository.GetById(id)
	return promotion
}

func GetRecords() []*model.Promotion {
	promotions := repository.GetAll()
	return promotions
}

func AddRecord(promotion model.Promotion) {
	repository.Add(promotion)
}
