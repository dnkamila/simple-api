package database

import (
	"gopkg.in/pg.v5"
	"net"
	"os"
)

type DB struct {
	*pg.DB
}

var db DB

func init() {
	db = DB{pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Addr:     net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})}
}

func GetDB() DB {
	return db
}
