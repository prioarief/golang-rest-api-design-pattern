package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prioarief/golang-rest-api-design-pattern/models"
	"github.com/prioarief/golang-rest-api-design-pattern/repositories"
)

type TodoController struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoController(todoRepository *repositories.TodoRepository) *TodoController {
	return &TodoController{todoRepository: todoRepository}
}

func (tc *TodoController) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.todoRepository.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) GetAll(c *gin.Context) {
	todos, err := tc.todoRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}
func (tc *TodoController) GetDetail(c *gin.Context) {
	id := c.Param("id")

	todo, err := tc.todoRepository.GetDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
