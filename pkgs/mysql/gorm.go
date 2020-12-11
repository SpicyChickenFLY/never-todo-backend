package mysql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lingdor/stackerror"
)

// GormDB is a connection pool of database
var GormDB *gorm.DB

// CreateGormConn create connection pool to mysql with gorm
func CreateGormConn(
	driverName,
	userName, userPwd,
	serverHost, serverPort,
	dbName, dbCharset string) error {
	var err error
	GormDB, err = gorm.Open(
		driverName,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			userName, userPwd,
			serverHost, serverPort,
			dbName, dbCharset))
	// defer GormDB.Close()
	if err != nil {
		return stackerror.New(err.Error())
	}
	log.Println("success to connect database")
	return nil
}

// CloseGormConn is a func to close connection pool
func CloseGormConn() {
	GormDB.Close()
}

// StopTransaction is a func to decide rollback or commit
func StopTransaction(tx *gorm.DB, err error) error {
	if err != nil {
		if err := tx.Rollback().Error; err != nil {
			return stackerror.New(err.Error())
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
