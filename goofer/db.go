*package goofer

import (
	"database/sql"
	"log"

	"github.com/gooferOrm/goofer/dialect"
	"github.com/gooferOrm/goofer/engine"
	"github.com/gooferOrm/goofer/schema"

	_ "github.com/mattn/go-sqlite3"
)
const DB_PATH = "./mockwails.db"
type GetClientReturn struct {
	Client *engine.Client
	db     *sql.DB
	Error  error
}
func GetClient(models ...schema.Entity) GetClientReturn {
	db, err := sql.Open("sqlite3", DB_PATH)
    if err != nil {
        log.Fatalf("open db: %v", err)
		return GetClientReturn{Error: err}
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
		return nil, err
    }
	return gooferClient, nil
}

func CreateServer(server Server) error {
	client, err := GetClient(&Server{})
	if err != nil {
		return err
	}
	serverRepo := engine.Repo[Server](client)
	return serverRepo.Save(&server)
}

func GetAllServers() ([]Server, error) {
	client, err := GetClient(&Server{})
	if err != nil {
		return nil, err
	}
	serverRepo := engine.Repo[Server](client)
	servers, err := serverRepo.Find().All()
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func UpdateServer(server Server) error {
	client, err := GetClient(&Server{})
	if err != nil {
		return err
	}
	serverRepo := engine.Repo[Server](client)
	return serverRepo.Save(&server)
}

func DeleteServer(ID uint) error {
	client, err := GetClient(&Server{})
	if err != nil {
		return err
	}
	serverRepo := engine.Repo[Server](client)
	return serverRepo.DeleteByID(ID)
}