package route

import (
	"github.com/SpicyChickenFLY/never-todo-backend/controller"
	todoCtrl "github.com/SpicyChickenFLY/never-todo-backend/controller/todo"
	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Group: Todo List
	groupTodo := router.Group("/api/v1/todo")
	{
		groupTodo.GET("/all", todoCtrl.GetAll)

		groupTask := groupTodo.Group("/task")
		{
			groupTask.GET("/", todoCtrl.GetAllTask)
			groupTask.POST("/", todoCtrl.AddTask)
			groupTask.DELETE("/:id", todoCtrl.DelTask)
			groupTask.PUT("/", todoCtrl.UpdTask)
		}
		groupTag := groupTodo.Group("/tag")
		{
			groupTag.GET("/", todoCtrl.GetAllTag)
			groupTag.POST("/", todoCtrl.AddTag)
			groupTag.DELETE("/:tag_id", todoCtrl.DelTag)
			groupTag.PUT("/", todoCtrl.UpdTag)
		}
		groupTaskTag := groupTodo.Group("/task-tag")
		{
			groupTaskTag.GET("/", todoCtrl.GetAllTaskTag)
			groupTaskTag.POST("/", todoCtrl.AddTaskTag)
			groupTaskTag.DELETE("/:id", todoCtrl.DelTaskTag)
		}
		groupFullTask := groupTodo.Group("/fulltask")
		{
			groupFullTask.GET("/", todoCtrl.GetAllFullTask)
			groupFullTask.GET("/content/:content", todoCtrl.GetFullTasksByContent)
			groupFullTask.GET("/tag/:tag_id", todoCtrl.GetFullTaskByTag)
			groupFullTask.POST("/", todoCtrl.AddFullTask)
			groupFullTask.DELETE("/:id", todoCtrl.DelFullTask)
			groupFullTask.PUT("/", todoCtrl.UpdFullTask)
		}
		groupSync := groupTodo.Group("/sync")
		{
			groupSync.GET("/", controller.GetLastSyncTime)
			groupSync.GET("/:sync_time", controller.SyncFromServer)
			groupSync.POST("/", controller.SyncToServer)
		}
		// groupUser := groupTodo.Group("/user")
		// {
		// 	groupUserInfo := groupUser.Group("/info")
		// 	{
		// 		// groupUserInfo.GET("/", controller.GetAllUserInfo)
		// 		groupUserInfo.GET("/:user_id", controller.GetUserInfoByID)
		// 		groupUserInfo.POST("/", controller.AddUserInfo)
		// 		groupUserInfo.POST("/:user_id", controller.DelUserInfo)
		// 		groupUserInfo.POST("/", controller.UpdUserInfo)
		// 	}

		// 	groupUser.POST("/login", controller.CheckUserLogin)
		// 	groupUser.POST("/logout", controller.CheckUserLogout)
		// }
	}
	router.GET("/", controller.ShowUI)
	return router
}
