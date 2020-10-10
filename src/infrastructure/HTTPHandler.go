package infrastructure

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/awesome-linus/go-gin-mysql-todo-api/src/application"
	"github.com/awesome-linus/go-gin-mysql-todo-api/src/domain"
	"github.com/awesome-linus/go-gin-mysql-todo-api/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler() *Handler {
	return &Handler{}
}

func NewHandlerWithMySQL(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) TodoList(c *gin.Context) {
	repo := &repository.TodoRepository{DB: h.DB}
	ms := application.TodoApplication{TodoRepository: repo}

	ml, err := ms.FetchAllFromMySQL()

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	if err != nil {
		CreateErrorResponse(c, err)
		return
	}

	CreateJsonResponse(c, http.StatusOK, ml.Todos)
}

func (h *Handler) ShowTodo(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("todoId"))

	repo := &repository.TodoRepository{DB: h.DB}
	ms := application.TodoApplication{TodoRepository: repo}

	req := &application.TodoFetchRequest{TodoID: todoID}
	todo, err := ms.FetchFromMySQL(*req)

	if err != nil {
		CreateErrorResponse(c, err)
		return
	}

	CreateJsonResponse(c, http.StatusOK, todo.Todo)
}

func (h *Handler) AddTodo(c *gin.Context) {
	todo := domain.Todo{}
	fmt.Println(todo)
	err := c.Bind(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	repo := &repository.TodoRepository{DB: h.DB}
	ms := application.TodoApplication{TodoRepository: repo}

	req := &application.TodoRegisterRequest{Todo: todo}
	createdTodo, err := ms.RegisterToMySQL(*req)

	if err != nil {
		CreateErrorResponse(c, err)
		return
	}

	CreateJsonResponse(c, http.StatusOK, createdTodo.Todo)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("todoId"))

	repo := &repository.TodoRepository{DB: h.DB}
	ms := application.TodoApplication{TodoRepository: repo}

	req := &application.TodoDeleteRequest{TodoID: todoID}
	todo, err := ms.DeleteFromMySQL(*req)

	if err != nil {
		CreateErrorResponse(c, err)
		return
	}

	CreateJsonResponse(c, http.StatusOK, todo.Todo)
}

func (h *Handler) ChangeTodo(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("todoId"))

	todo := domain.Todo{}
	fmt.Println(todo)
	err := c.Bind(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	repo := &repository.TodoRepository{DB: h.DB}
	ms := application.TodoApplication{TodoRepository: repo}

	req := &application.TodoUpdateRequest{TodoID: todoID, Todo: todo}
	updatedTodo, err := ms.UpdateToMySQL(*req)

	if err != nil {
		CreateErrorResponse(c, err)
		return
	}

	CreateJsonResponse(c, http.StatusOK, updatedTodo.Todo)
}
