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
	createStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (identifier integer generated always as identity, id char(50), price real, expirationDate char(50))", TABLE)
	_, err := conn.DB.Exec(createStatement)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func DropTable() {
	conn = NewConnection()
	dropStatement := fmt.Sprintf("DROP TABLE IF EXISTS %s", TABLE)
	_, err := conn.DB.Exec(dropStatement)
	if err != nil {
		log.Fatal(err)
	}
	conn = nil
}

func Add(p model.Promotion) {
	conn := getConnection()
	insertStatement := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1, $2, $3)", TABLE)
	conn.Insert(insertStatement, p.Id, p.Price, p.ExpirationDate)
}

func AddAll(promotions []*model.Promotion) {
	conn := getConnection()

	var (
		placeholders []string
		values       []interface{}
	)

	for index, promotion := range promotions {
		placeholders = append(placeholders, fmt.Sprintf("(DEFAULT,$%d,$%d,$%d)",
			index*3+1,
			index*3+2,
			index*3+3,
		))

		values = append(values, promotion.Id, promotion.Price, promotion.ExpirationDate)
	}

	conn.BatchInsert(TABLE, placeholders, values)
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

func GetAll(offset any, limit any) []*model.Promotion {
	promotions := make([]*model.Promotion, 0)
	conn := getConnection()
	statement := fmt.Sprintf("SELECT * FROM %s LIMIT %s OFFSET %s", TABLE, limit, offset)
	rows := conn.SelectAll(statement)
	for rows.Next() {
		id := ""
		promotion := model.Promotion{}
		err := rows.Scan(&id, &promotion.Id, &promotion.Price, &promotion.ExpirationDate)
		if err != nil {
			log.Fatal(err)
		}
		promotions = append(promotions, &promotion)
	}
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	return promotions
}
