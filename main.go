package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/mysql"
	"github.com/SpicyChickenFLY/never-todo-backend/route"

	_ "github.com/go-sql-driver/mysql"
)

const ( // MYSQL CONFIG
	mysqlDriverName      = "mysql"
	mysqlServerHost      = "localhost"
	mysqlServerPort      = "3306"
	mysqlUserName        = "root"
	mysqlUserPwd         = "123"
	mysqlDatabaseName    = "never_todo"
	mysqlDatabaseCharset = "utf8"
)

const ( // GIN CONFIG
	port = ":8080"
)

func main() {
	// get mysql root@localhost password
	userPwd := ""
	fmt.Printf("Please enter password for mysql user root@localhost: ")
	fmt.Scanln(&userPwd)
	if userPwd == "" {
		userPwd = mysqlUserPwd
	}

	// Initialize MySQL connection
	mysql.CreateGormConn(
		mysqlUserName, userPwd,
		mysqlServerHost, mysqlServerPort,
		mysqlDatabaseName, mysqlDatabaseCharset)

	server := &http.Server{
		Addr:    port,
		Handler: route.InitRouter(),
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("server encount error while listen and serve:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}
