package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prioarief/golang-rest-api-design-pattern/helpers"
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
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	if err := tc.todoRepository.Create(&todo); err != nil {
		helpers.Response(c, 0, err.Error(), nil, 500)
		return
	}

	helpers.Response(c, 1, "Success Add To Do", todo, 201)
}

func (tc *TodoController) GetAll(c *gin.Context) {
	todos, err := tc.todoRepository.GetAll()
	if err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	helpers.Response(c, 1, "To Do List", todos, 200)
}
func (tc *TodoController) GetDetail(c *gin.Context) {
	id := c.Param("id")

	todo, err := tc.todoRepository.GetDetail(id)
	if err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	helpers.Response(c, 1, "To Do", todo, 200)
}

func (tc *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := tc.todoRepository.GetDetail(id)
	if err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	if err2 := tc.todoRepository.Delete(id); err2 != nil {
		helpers.Response(c, 0, err2.Error(), nil, 400)
		return
	}

	helpers.Response(c, 1, "Success Delete To Do", nil, 200)
}

func (tc *TodoController) Update(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	_, err := tc.todoRepository.GetDetail(id)
	if err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helpers.PanicHelper("Error parsing integer", c)
		return
	}

	if err := c.ShouldBind(&todo); err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	todo.ID = int(num)

	if err := tc.todoRepository.Update(&todo); err != nil {
		helpers.Response(c, 0, err.Error(), nil, 400)
		return
	}

	helpers.Response(c, 1, "Success Update To Do", nil, 200)
}
