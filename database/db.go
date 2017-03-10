package database

import (
	"gopkg.in/pg.v5"
	"net"
	"os"
	"simple-api/database/constants"
)

type DB struct {
	*pg.DB
}

var (
	db         DB
	dbUsername = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbHostname = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
)

func init() {
	setupDBOptions()
	db = DB{pg.Connect(&pg.Options{
		User:     dbUsername,
		Password: dbPassword,
		Database: dbName,
		Addr:     net.JoinHostPort(dbHostname, dbPort),
	})}
}

func setupDBOptions() {
	if dbUsername == "" {
		dbUsername = constants.DB_USERNAME
	}
	if dbPassword == "" {
		dbPassword = constants.DB_PASSWORD
	}
	if dbName == "" {
		dbName = constants.DB_NAME
	}
	if dbHostname == "" {
		dbHostname = constants.DB_HOST
	}
	if dbPort == "" {
		dbPort = constants.DB_PORT
	}
}

func GetDB() DB {
	return db
}
