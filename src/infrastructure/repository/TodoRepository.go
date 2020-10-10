package repository

import (
	"errors"
	"log"

	"github.com/awesome-linus/go-gin-mysql-todo-api/src/domain"
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (m *TodoRepository) FindAll() (domain.Todos, error) {
	var todos domain.Todos
	m.DB.Order("ID asc").Find(&todos)
	return todos, nil
}

func (m *TodoRepository) Find(todoID int) (domain.Todo, error) {
	var todo domain.Todo
	isRecordNotFound := m.DB.First(&todo, todoID).RecordNotFound()

	if isRecordNotFound {
		// TODO: Add Error Custom Handling
		return todo, errors.New("エラーだよ～")
	}

	return todo, nil
}

func (m *TodoRepository) Register(todo domain.Todo) (domain.Todo, error) {
	result := m.DB.Create(&todo)

	if result.Error != nil {
		panic("Error")
	}

	if result.RowsAffected != 1 {
		log.Printf("Found Unexpected Affected Rows: %d", result.RowsAffected)
	}

	var createdTodo domain.Todo
	m.DB.First(&createdTodo, todo.ID)

	return createdTodo, nil
}

func (m *TodoRepository) Delete(todoID int) (domain.Todo, error) {
	var todo domain.Todo
	result := m.DB.Delete(&todo, todoID)

	if result.Error != nil {
		// TODO: Add Error Custom Handling
		return todo, errors.New("エラーだよ～")
	}

	if result.RowsAffected != 1 {
		log.Printf("Found Unexpected Affected Rows: %d", result.RowsAffected)
		// TODO: Add Error Custom Handling
		return todo, errors.New("エラーだよ～")
	}

	return todo, nil
}

func (m *TodoRepository) Update(todoID int, todo domain.Todo) (domain.Todo, error) {
	result := m.DB.Model(&todo).Where("ID = ?", todoID).Update(todo)

	if result.Error != nil {
		// TODO: Add Error Custom Handling
		return todo, errors.New("エラーだよ～")
	}

	if result.RowsAffected != 1 {
		log.Printf("Found Unexpected Affected Rows: %d", result.RowsAffected)
		// TODO: Add Error Custom Handling
		return todo, errors.New("エラーだよ～")
	}

	var updatedTodo domain.Todo
	m.DB.First(&updatedTodo, todoID)

	return updatedTodo, nil
}
