package repository

import (
	"database/sql"
	"example/promote/configuration"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewConnection() *Database {
	c, err := configuration.New()
	if err != nil {
		log.Panic("Configuration Error", err)
	}
	connectionDetails := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.DB.USER, c.DB.PASSWORD, c.DB.NAME, c.DB.HOST, c.DB.PORT)
	db, err := sql.Open(c.DB.DIALECT, connectionDetails)
	if err != nil {
		log.Fatal(err)
	}
	database := &Database{}
	database.DB = db
	return database
}

func (database *Database) SelectAll(statement string) *sql.Rows {
	rows, err := database.DB.Query(statement)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func (database *Database) Select(statement string, id int) *sql.Row {
	row := database.DB.QueryRow(statement, strconv.Itoa(id))
	return row
}

func (database *Database) Insert(statement string, id string, price float32, date string) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	_, err = stmt.Exec(id, price, date)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
