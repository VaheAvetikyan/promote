package repository

import (
	"database/sql"
	"example/promote/configuration"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	"time"
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
	db.SetMaxOpenConns(c.DB.PROPERTIES.MaxOpenConnections)
	db.SetMaxIdleConns(c.DB.PROPERTIES.MaxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(c.DB.PROPERTIES.ConnMaxIdleTime))

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

func (database *Database) Insert(statement string, id string, price float64, date string) {
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

func (database *Database) BatchInsert(table string, placeholders []string, values []interface{}) {
	go func() {
		txn, err := database.DB.Begin()
		if err != nil {
			log.Fatal("Could not start a new transaction. ", err)
		}

		insertStatement := fmt.Sprintf("INSERT INTO %s VALUES %s", table, strings.Join(placeholders, ","))
		stmt, err := txn.Prepare(insertStatement)
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			txn.Rollback()
			log.Fatal("Failed to insert multiple records at once.", err)
		}

		if err := txn.Commit(); err != nil {
			log.Fatal("Failed to commit transaction.", err)
		}

		if err = stmt.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
