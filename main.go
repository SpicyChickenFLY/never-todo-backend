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

	"github.com/gin-gonic/gin"
	"spicychicken.top/NeverTODO/backend/controller"
	"spicychicken.top/NeverTODO/backend/pkgs/middleware"
	"spicychicken.top/NeverTODO/backend/pkgs/mysql"

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
	GIN_PORT = ":8080"
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
		mysqlDriverName,
		mysqlUserName, userPwd,
		mysqlServerHost, mysqlServerPort,
		mysqlDatabaseName, mysqlDatabaseCharset)

	router := gin.Default()
	router.Use(middleware.Cors())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Group: Todo List
	groupTodo := router.Group("/todo")
	{
		groupTodo.GET("/all", controller.GetAll)
		groupTask := groupTodo.Group("/task")
		{
			groupTask.POST("/add", controller.AddNewTask)
			groupTask.POST("/del", controller.DelOldTask)
		}
		groupTag := groupTodo.Group("/tag")
		{
			groupTag.POST("/add", controller.AddNewTag)
			groupTag.POST("/del", controller.DelOldTag)
		}
		groupTaskList := groupTodo.Group("/tasklist")
		{
			groupTaskList.GET("/all", controller.GetTaskList)
			groupTaskList.GET("/tag", controller.GetTaskListByTag)
		}

	}
	groupRequest := router.Group("/request")
	{
		groupRequest.GET("/ui", controller.ShowUI)
		groupRequest.GET("/send", controller.SendRequest)
		groupSim := groupRequest.Group("/sim")
		{
			groupSimTask := groupSim.Group("/task")
			{
				groupSimTask.GET("/add", controller.SimAddTask)
				groupSimTask.GET("/del", controller.SimDelTask)
			}
			groupSimTag := groupSim.Group("/tag")
			{
				groupSimTag.GET("/add", controller.SimAddTag)
				groupSimTag.GET("/del", controller.SimDelTag)
			}
		}
	}

	server := &http.Server{
		Addr:    GIN_PORT,
		Handler: router,
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("server encount error while listen and serve")
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
	mysql.CloseGormConn()
	log.Println("Server exiting")
}
