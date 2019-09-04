package dbService

import (
	"errors"
	"github.com/wangjiaming0909/GoP/service"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type DBService struct {
	service.Service
	db *sql.DB
	config *DBServiceConfig
}

type DBServiceConfig struct {
	serviceName string
	dbDriver string
	userName string
	password string
	dbName string
	address string
	port uint
	//...
}

func (c *DBServiceConfig) GetDSN() string {
	return ""
}

func NewDBServiceConfig(driver string, uN string, pword string, dbN string, addr string, p uint) DBServiceConfig {
	return DBServiceConfig{
		serviceName: "DBService",
		dbDriver: driver,
		userName: uN,
		password: pword,
		dbName: dbN,
		address: addr,
		port: p,
	}
}

func NewDBService(config *DBServiceConfig) DBService {
	return DBService{
			config: config,
			db: nil,
		}
}

func (t *DBService) Start() error{
	db, err := sql.Open(t.config.dbDriver, t.config.GetDSN())
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	t.Status = service.Started
	return nil
}

func (t *DBService) Stop() error {
	if t.db == nil {
		return errors.New("not connected")
	}
	err := t.db.Close()
	t.Status = service.Stopped
	return err
}

