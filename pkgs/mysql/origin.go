package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lingdor/stackerror"
)

// OriginDB is a connection pool of database
var OriginDB *sql.DB

// CreateOriginConn is a func to create a connection pool
func CreateOriginConn(
	driverName,
	userName, userPwd,
	serverHost, serverPort,
	dbName, dbCharset string) error {
	// connect the Mysql instance and select specified db
	var err error
	OriginDB, err = sql.Open(
		driverName,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s",
			userName, userPwd,
			serverHost, serverPort,
			dbName, dbCharset))
	if err != nil {
		return stackerror.New(err.Error())
	}
	log.Println("success to connect database")
	if err := OriginDB.Ping(); err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// CloseOriginConn is a func to close connection pool
func CloseOriginConn() {
	OriginDB.Close()
}
