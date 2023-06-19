package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prioarief/golang-rest-api-design-pattern/controllers"
	"github.com/prioarief/golang-rest-api-design-pattern/repositories"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// Create repository and controller instances
	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("/", todoController.GetAll)
		}
	}

	return r
}
