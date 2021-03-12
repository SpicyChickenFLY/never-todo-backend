package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/romberli/log"
	"gopkg.in/ini.v1"

	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/mysql"
	"github.com/SpicyChickenFLY/never-todo-backend/route"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultLogFileRelPath  = "/log/never-todo.log"
	defaultConfFileRelPath = "/config/never-todo.ini"
)

const ( // GIN CONFIG
	port = ":8080"
)

func main() {
	// currDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("get current Directory failed: %s\n", err.Error())
	// 	panic(err)
	// }
	// currDir = strings.ReplaceAll(currDir, "\\", "/")

	currDir := "/mnt/d/Code/go/src/github.com/SpicyChickenFLY/never-todo-backend"

	// Initialize log
	if _, _, err := log.InitLoggerWithDefaultConfig(
		path.Join(currDir, defaultLogFileRelPath)); err != nil {
		fmt.Printf("Init logger failed: %s\n", err.Error())
		panic(err)
	}

	log.Info("=============================")
	log.Info("Program Started")

	ginMode := flag.String("m", "debug", "GIN_MODE:debug/release/test")
	gin.SetMode(*ginMode)
	configFile := flag.String("c", path.Join(currDir, defaultConfFileRelPath), "configure file")
	cfg, err := ini.Load(*configFile)
	if err != nil {
		log.Error(err.Error())
		log.Info("=============================")
		panic(err)
	}

	// get mysql root@localhost password
	dbType := cfg.Section("db").Key("type").String()
	if dbType == "mysql" {
		serverHost := cfg.Section("db").Key("server_host").String()
		serverPort := cfg.Section("db").Key("server_port").String()
		userName := cfg.Section("db").Key("user_name").String()
		userPwd := cfg.Section("db").Key("user_pwd").String()
		dbName := cfg.Section("db").Key("db_name").String()
		dbCharset := cfg.Section("db").Key("db_charset").String()
		// Initialize MySQL connection
		if err := mysql.CreateGormConn(
			userName, userPwd,
			serverHost, serverPort,
			dbName, dbCharset); err != nil {
			log.Error(err.Error())
			log.Info("=============================")
			panic(err)
		}
		log.Info("mysql initialization compelete")
	}

	server := &http.Server{
		Addr:    port,
		Handler: route.InitRouter(),
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("encounter error while listen and serve:\n", err)
			log.Info("=============================")
			panic(err)
		}
		log.Info("server initialization compelete")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("Server Shutdown:\n", err.Error())
		log.Info("=============================")
		panic(err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 1 seconds.")
	}
	log.Info("Server exiting")
}
