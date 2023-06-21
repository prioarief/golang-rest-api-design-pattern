package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prioarief/golang-rest-api-design-pattern/configs"
	"github.com/prioarief/golang-rest-api-design-pattern/controllers"
	"github.com/prioarief/golang-rest-api-design-pattern/repositories"
)

func TodoRoutes(r *gin.RouterGroup) {
	db := configs.Database()

	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	todos := r.Group("/todos")
	{
		todos.GET("/", todoController.GetAll)
		todos.GET("/:id", todoController.GetDetail)
	}
}
