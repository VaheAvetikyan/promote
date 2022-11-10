package repository

import (
	"database/sql"
	"example/promote/model"
	"log"
)

func Add(promotion model.Promotion) {

}

func GetAll() []*model.Promotion {
	promotions := make([]*model.Promotion, 0)
	promotion := model.Promotion{}
	conn := NewConnection()
	rows := conn.SelectAll("promotion")
	for rows.Next() {
		err := rows.Scan(&promotion.Id, &promotion.Price, &promotion.ExpirationDate)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("retrieved promotion with id: ", promotion.Id)
		promotions = append(promotions, &promotion)
	}
	return promotions
}

func GetById(id int) model.Promotion {
	promotion := model.Promotion{}
	conn := NewConnection()
	row := conn.Select("promotion", id)
	err := row.Scan(&promotion.Id, &promotion.Price, &promotion.ExpirationDate)
	if err != sql.ErrNoRows {
		log.Fatal(err)
	}
	log.Println("retrieved promotion with id: ", id)
	return promotion
}
