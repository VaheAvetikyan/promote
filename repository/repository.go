package repository

import (
	"database/sql"
	"example/promote/configuration"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewConnection() *Database {
	configuration, err := configuration.New()
	if err != nil {
		log.Panic("Configuration Error", err)
	}
	db, err := sql.Open(configuration.DB.DIALECT, "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}
	database := &Database{}
	database.DB = db
	return database
}

func (database *Database) SelectAll(table string) *sql.Rows {
	rows, err := database.DB.Query("SELECT * FROM " + table)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows
}

func (database *Database) Select(table string, id int) *sql.Row {
	statement := "SELECT * FROM " + table + "WHERE ID = $1"
	row := database.DB.QueryRow(statement, table, strconv.Itoa(id))
	return row
}

func (database *Database) Insert(table string, id string, price float32, date string) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO " + table + "VALUES ($1, $2, $3)")
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
