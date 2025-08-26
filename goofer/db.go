package goofer

import (
	"database/sql"
	"log"

	"github.com/gooferOrm/goofer/dialect"
	"github.com/gooferOrm/goofer/engine"
	"github.com/gooferOrm/goofer/schema"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "./mockwails.db"

type DBClient struct {
	Client *engine.Client
	DB     *sql.DB
	Err    error
}

func GetClient(models ...schema.Entity) DBClient {
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatalf("open db: %v", err)
		return DBClient{Err: err}
	}
	// defer db.Close()

	// engine setup:
	gooferClient, err := engine.NewClient(
		db,
		dialect.NewSQLiteDialect(),
		models...,
	)
	if err != nil {
		log.Fatalf("engine init: %v", err)
		return DBClient{Err: err}
	}
	return DBClient{Client: gooferClient, DB: db}
}

func CreateServer(server ServerEntity) error {
	dbclient := GetClient(&ServerEntity{})
	if dbclient.Err != nil {
		return dbclient.Err
	}
	defer dbclient.DB.Close()
	serverRepo := engine.Repo[ServerEntity](dbclient.Client)
	return serverRepo.Save(&server)
}

func GetAllServers() ([]ServerEntity, error) {
	dbclient := GetClient(&ServerEntity{})
	if dbclient.Err != nil {
		return nil, dbclient.Err
	}
	defer dbclient.DB.Close()
	serverRepo := engine.Repo[ServerEntity](dbclient.Client)
	servers, err := serverRepo.Find().All()
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func UpdateServer(server ServerEntity) error {
	dbclient := GetClient(&ServerEntity{})
	if dbclient.Err != nil {
		return dbclient.Err
	}
	defer dbclient.DB.Close()
	serverRepo := engine.Repo[ServerEntity](dbclient.Client)
	return serverRepo.Save(&server)
}

func DeleteServer(ID uint) error {
	dbclient := GetClient(&ServerEntity{})
	if dbclient.Err != nil {
		return dbclient.Err
	}
	defer dbclient.DB.Close()
	serverRepo := engine.Repo[ServerEntity](dbclient.Client)
	return serverRepo.DeleteByID(ID)
}
