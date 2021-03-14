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
	router.LoadHTMLGlob("templates/*/**")
	router.Static("/static", "./static")

	// Group: Todo List
	groupTodo := router.Group("/todo")
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
			groupTag.DELETE("/:tag_id", controller.DelOldTag)
			groupTag.PUT("/", controller.UpdOldTag)
		}
		groupFullTask := groupTodo.Group("/fulltask")
		{
			groupFullTask.GET("/", controller.GetAllFullTask)
			groupFullTask.GET("/:tag_id", controller.GetFullTaskByTag)
			groupFullTask.POST("/", controller.AddNewFullTask)
			groupFullTask.DELETE("/:task_id", controller.DelOldFullTask)
			groupFullTask.PUT("/", controller.UpdOldFullTask)
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
