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
			groupTag.POST("/", controller.AddNewTag)
			groupTag.DELETE("/:id", controller.DelOldTag)
			groupTag.PUT("/:id", controller.UpdOldTag)
		}
		groupFullTask := groupTodo.Group("/fulltask")
		{
			groupFullTask.GET("/", controller.GetAllFullTask)
			groupFullTask.GET("/content/:content", controller.GetFullTasksByContent)
			groupFullTask.GET("/tag/:id", controller.GetFullTaskByTag)
			groupFullTask.POST("/", controller.AddNewFullTask)
			groupFullTask.DELETE("/:id", controller.DelOldFullTask)
			groupFullTask.PUT("/:id", controller.UpdOldFullTask)
		}
		groupSync := groupTodo.Group("/sync")
		{
			groupSync.GET("/", controller.GetLastSyncTime)
			groupSync.GET("/:sync_time", controller.SyncFromServer)
			groupSync.POST("/", controller.SyncToServer)
		}
	}
	router.GET("/", controller.ShowUI)
	return router
}
