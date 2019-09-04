package dbService

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func SqlDaoService() {
	db, err := sql.Open("mssql", "server=tcp:127.0.0.1:49172;database=userdb;integrated security=SSPI")
	checkError(err)

	db.Query("drop database if exists userdb")
	db.Query("create database userdb")
	db.Query("use userdb")

	db.Query("create table userdb.userInfo(userId int(11) PRIMARY KEY, name varchar(20), email varchar(20), phone varchar(20), password varchar(20))")
	db.Query("insert into userdb.userInfo values(101, '姓名1', 'address1', '18811110000'), (102, '姓名2', 'address2', '18811220000')")

}

func checkError(errMsg error) {
	if errMsg != nil {
		log.Fatal(errMsg)
	}
}
