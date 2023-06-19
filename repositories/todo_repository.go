package repositories

import (
	"github.com/prioarief/golang-rest-api-design-pattern/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (tr *TodoRepository) Create(todo *models.Todo) error {
	return tr.db.Create(todo).Error
}

func (tr *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := tr.db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}
