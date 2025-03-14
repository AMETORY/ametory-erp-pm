package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	projectHandler := handlers.NewProjectHandler(erpContext)
	taskHandler := handlers.NewTaskHandler(erpContext)

	projectGroup := r.Group("/project")
	projectGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		projectGroup.GET("/list", projectHandler.GetProjectsHandler)
		projectGroup.GET("/templates", projectHandler.GetTemplatesHandler)
		projectGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:project:create"}), projectHandler.CreateProjectHandler)
		projectGroup.GET("/:id", projectHandler.GetProjectHandler)
		projectGroup.GET("/:id/members", projectHandler.GetMembersHandler)
		projectGroup.POST("/:id/add-member", projectHandler.AddMemberHandler)
		projectGroup.PUT("/:id/update-column", projectHandler.UpdateColumnHandler)
		projectGroup.PUT("/:id/rearrange-columns", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:project:update"}), projectHandler.RearrangeColumnsHandler)
		projectGroup.PUT("/:id/add-column", projectHandler.AddNewColumnHandler)
		projectGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:project:update"}), projectHandler.UpdateProjectHandler)
		projectGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:project:delete"}), projectHandler.DeleteProjectHandler)
		projectTaskGroup := projectGroup.Group("/:id/task")
		{
			projectTaskGroup.GET("/:taskId/detail", taskHandler.GetTaskDetailHandler)
			projectTaskGroup.PUT("/rearrange", taskHandler.RearrangeTaskHandler)
			projectTaskGroup.GET("/list", taskHandler.GetTasksHandler)
			projectTaskGroup.POST("/create", taskHandler.CreateTaskHandler)
			projectTaskGroup.PUT("/:taskId/move", taskHandler.MoveTaskHandler)
			projectTaskGroup.PUT("/:taskId/update", taskHandler.UpdateTaskHandler)
			projectTaskGroup.POST("/:taskId/comment", taskHandler.AddCommentHandler)
			projectTaskGroup.PUT("/:taskId/add-plugin", taskHandler.AddPluginHandler)
			projectTaskGroup.GET("/:taskId/get-plugins", taskHandler.GetTaskPluginsHandler)
			projectTaskGroup.GET("/:taskId/plugin/:pluginId", taskHandler.GetDataPluginHandler)
			// taskGroup.GET("/:id", taskHandler.GetTaskHandler)
			// taskGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:task:update"}), taskHandler.UpdateTaskHandler)
			// taskGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"project_management:task:delete"}), taskHandler.DeleteTaskHandler)
		}
	}

	taskGroup := r.Group("/task")
	taskGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		taskGroup.GET("/my", taskHandler.MyTaskHandler)
		taskGroup.GET("/watched", taskHandler.WatchedTaskHandler)
	}

}
