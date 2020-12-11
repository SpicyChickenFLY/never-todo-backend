package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lingdor/stackerror"
)

const (
	// TASK - todo_task
	SQL_ADD_TASK = `
		INSERT INTO todo_task(content) 
			VALUES("%s")`
	SQL_GET_ALL_TASKS = `
		SELECT * FROM todo_task`
	SQL_GET_TASKS_BY_ID = `
		SELECT * FROM todo_task
			WHERE id=%d`
	SQL_GET_TASKS_NUM_BY_ID = `
		SELECT COUNT(*) FROM todo_task
			WHERE id=%d`
	SQL_DEL_TASK_BY_ID = `
		DELETE FROM todo_task
			WHERE id=%d`
	// TAG - todo_tag
	SQL_ADD_TAG = `
		INSERT INTO todo_tag(content, description)
			VALUES("%s", "%s")`
	SQL_GET_ALL_TAGS = `
		SELECT * FROM todo_tag`
	SQL_GET_TAGS_BY_ID = `
		SELECT * FROM todo_tag
			WHERE id=%d`
	SQL_GET_TAGS_NUM_BY_ID = `
		SELECT COUNT(*) FROM todo_tag 
			WHERE id=%d`
	SQL_GET_TAGS_NUM_BY_CONTENT = `
		SELECT COUNT(*) FROM todo_tag 
			WHERE content="%s"`
	SQL_DEL_TAG_BY_ID = `
		DELETE FROM todo_tag
			WHERE id=%d`
	// TASK_TAG - todo_task_tag
	SQL_ADD_TASK_TAG = `
		INSERT INTO todo_task_tag(task_id, tag_id)
			VALUES(%d, %d)`
	SQL_GET_ALL_TASK_TAG = `
		SELECT * FROM todo_task_tag`
	SQL_GET_TAGS_OF_TASK = `
		SELECT todo_tag.* FROM todo_tag
			INNER JOIN todo_task_tag
			ON todo_tag.id=todo_task_tag.tag_id
			AND task_id=%d;`
	SQL_GET_TASKS_BY_TAG = `
		SELECT todo_task.* FROM todo_task_tag 
			INNER JOIN todo_task
			ON todo_task.id=todo_task_tag.task_id
			WHERE tag_id=%s`
	SQL_DEL_TAGS_OF_TASK = `
		DELETE FROM todo_task_tag
			WHERE task_id=%d`
	SQL_DEL_TAG_OF_TASKS = `
		DELETE FROM todo_task_tag
			WHERE tag_id=%d`
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
