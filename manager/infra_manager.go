package manager

import (
	"database/sql"
	"fmt"
	"log"

	"Merchant-Bank/config"
)

type manager interface {
	DbConn() *sql.DB
}

type InfraManager struct {
	db  *sql.DB
	cfg config.AppConfig
}

func (i *InfraManager) dbInit() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name, i.cfg.SslMode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application filed to run", err)
			db.Close()
		}
	}()

	i.db = db
	fmt.Println("DB Connected")
}

func (i *InfraManager) DbConn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg config.AppConfig) manager {
	infra := InfraManager{
		cfg: cfg,
	}
	infra.dbInit()
	return &infra
}
