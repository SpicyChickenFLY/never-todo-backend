package route

import (
	"github.com/SpicyChickenFLY/never-todo-backend/controller"
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
		groupTodo.GET("/all", controller.GetAll)
		groupTask := groupTodo.Group("/task")
		{
			groupTask.GET("/", controller.GetAllTask)
		}
		groupTag := groupTodo.Group("/tag")
		{
			groupTag.GET("/", controller.GetAllTag)
			groupTag.POST("/", controller.AddTag)
			groupTag.DELETE("/:tag_id", controller.DelTag)
			groupTag.PUT("/", controller.UpdTag)
		}
		groupFullTask := groupTodo.Group("/fulltask")
		{
			groupFullTask.GET("/", controller.GetAllFullTask)
			groupFullTask.GET("/content/:content", controller.GetFullTasksByContent)
			groupFullTask.GET("/tag/:tag_id", controller.GetFullTaskByTag)
			groupFullTask.POST("/", controller.AddFullTask)
			groupFullTask.DELETE("/:task_id", controller.DelFullTask)
			groupFullTask.PUT("/", controller.UpdFullTask)
		}
		groupSync := groupTodo.Group("/sync")
		{
			groupSync.GET("/", controller.GetLastSyncTime)
			groupSync.GET("/:sync_time", controller.SyncFromServer)
			groupSync.POST("/", controller.SyncToServer)
		}
		groupUser := groupTodo.Group("/user")
		{
			groupUserInfo := groupUser.Group("/info")
			{
				// groupUserInfo.GET("/", controller.GetAllUserInfo)
				groupUserInfo.GET("/:user_id", controller.GetUserInfoByID)
				groupUserInfo.POST("/", controller.AddUserInfo)
				groupUserInfo.POST("/:user_id", controller.DelUserInfo)
				groupUserInfo.POST("/", controller.UpdUserInfo)
			}

			groupUser.POST("/login", controller.CheckUserLogin)
			groupUser.POST("/logout", controller.CheckUserLogout)
		}
	}
	router.GET("/", controller.ShowUI)
	return router
}
