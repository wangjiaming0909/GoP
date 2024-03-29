package dbservice

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjiaming0909/GoP/service"
)

//DBService DBService
type DBService struct {
	service.Service
	db     *sql.DB
	config *DBServiceConfig
}

//DBServiceConfig config of DB service
type DBServiceConfig struct {
	serviceName string
	dbDriver    string
	userName    string
	password    string
	dbName      string
	address     string
	port        uint
	//...
}

//GetDSN get sql dsn
func (c *DBServiceConfig) GetDSN() string {
	// return c.userName + ":" + c.password + "@tcp(" + c.address + ":" + string(c.port) + ")/user" + "?charset=utf8"
	return c.userName + ":" + c.password + "@/" + "user"
}

//NewDBServiceConfig new a DB service config
func NewDBServiceConfig(driver string, uN string, pword string, dbN string, addr string, p uint) DBServiceConfig {
	return DBServiceConfig{
		serviceName: "DBService",
		dbDriver:    driver,
		userName:    uN,
		password:    pword,
		dbName:      dbN,
		address:     addr,
		port:        p,
	}
}

//NewDBService new a service of DB
func NewDBService(config *DBServiceConfig) DBService {
	return DBService{
		config: config,
		db:     nil,
	}
}

//Start start the DB service
func (t *DBService) Start() error {
	db, err := sql.Open(t.config.dbDriver, t.config.GetDSN())
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	t.Status = service.Started
	Create_Table(db)
	return nil
}

//Stop stop the DB service
func (t *DBService) Stop() error {
	if t.db == nil {
		return errors.New("not connected")
	}
	err := t.db.Close()
	t.Status = service.Stopped
	return err
}

func Create_Table(db *sql.DB) {
	if db == nil {
		return
	}

	sql := `
	CREATE TABLE ` + `UserInfo` + `(
		uid INT(10) NOT NULL AUTO_INCREMENT,
		username VARCHAR(64) NULL DEFAULT NULL,
		password VARCHAR(64) NULL DEFAULT NULL,
		email VARCHAR(64) NULL DEFAULT NULL,
		created DATE NULL DEFAULT NULL,
		PRIMARY KEY(uid)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
