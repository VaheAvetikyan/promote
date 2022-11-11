package repository

import (
	"example/promote/model"
	"fmt"
	"log"
)

const TABLE string = "promotions"

var conn *Database

func getConnection() *Database {
	if conn != nil {
		return conn
	}
	conn = NewConnection()
	_, err := conn.DB.Exec("CREATE TABLE IF NOT EXISTS promotions (identifier integer generated always as identity, id char(50), price real, expirationDate char(50))")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func Add(p model.Promotion) {
	conn := getConnection()
	statement := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1, $2, $3)", TABLE)
	conn.Insert(statement, p.Id, p.Price, p.ExpirationDate)
}

func GetAll() []*model.Promotion {
	promotions := make([]*model.Promotion, 0)
	promotion := model.Promotion{}
	conn := getConnection()
	statement := fmt.Sprintf("SELECT * FROM %s", TABLE)
	rows := conn.SelectAll(statement)
	for rows.Next() {
		id := ""
		err := rows.Scan(&id, &promotion.Id, &promotion.Price, &promotion.ExpirationDate)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Retrieved promotion with id: ", promotion.Id)
		promotions = append(promotions, &promotion)
	}
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	return promotions
}

func GetById(id int) model.Promotion {
	promotion := model.Promotion{}
	conn := getConnection()
	statement := fmt.Sprintf("SELECT * FROM %s WHERE identifier=$1", TABLE)
	row := conn.Select(statement, id)
	err := row.Scan(&id, &promotion.Id, &promotion.Price, &promotion.ExpirationDate)
	if err != nil {
		log.Print(err)
	} else {
		log.Print("Retrieved promotion with id: ", id)
	}
	return promotion
}
