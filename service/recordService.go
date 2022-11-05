package service

import (
	"example/promote/model"
	"example/promote/repository"
)

func GetRecord(id int) model.Record {
	record := repository.New().GetById(id)
	return record
}

func GetRecords() []model.Record {
	records := repository.New().GetAll()
	return records
}
