package goofer

import (
	"database/sql"
	"log"

	"github.com/bluesky-social/indigo/models"
	"github.com/gooferOrm/goofer/dialect"
	"github.com/gooferOrm/goofer/engine"
)
const DB_PATH = "./mockwails.db"

func GetClient(models ...interface{}) (*engine.Client, error) {
	db, err := sql.Open("sqlite3", DB_PATH)
    if err != nil {
        log.Fatalf("open db: %v", err)
		return nil, err
    }
    defer db.Close()

    // engine setup:
    gooferClient, err := engine.NewClient(
        db,
        dialect.NewSQLiteDialect(), 
        models...,
    )
    if err != nil {
        log.Fatalf("engine init: %v", err)
		return nil, err
    }
	return gooferClient, nil
}

func CreateServer() error {
	// TODO: Implementation for creating a server in the database
	return nil
}