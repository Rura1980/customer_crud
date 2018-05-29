package app

import (
    "database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var appDb *DB

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "customer_db_user"
	PASSWORD = "12345678"
	DB_NAME  = "customer_db"
)

type DB struct {
	*sql.DB
}

type rowScanner interface {
	Scan(target ...interface{}) error
}

// return database reference for a data source
func Connect() (error) {
	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DB_NAME, PORT)
	db, err := sql.Open("postgres", dataSource) // ignoring error since no actual connection is establish
	if err != nil {
		return err
	}
	err = db.Ping()                          // here comes actual connection tryout
	if err != nil {
		return err
	}
	
	appDb = &DB{db}
	return nil
}

func Disconnect() {
	appDb.Close()
}

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{s, true}
}

func NullString2String(s sql.NullString ) string {
	if !s.Valid {
		return ""
	}
	return s.String
}
