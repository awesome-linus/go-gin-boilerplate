package application

import (
	"github.com/awesome-linus/go-gin-mysql-todo-api/src/domain"
)

type TodoApplication struct {
	TodoRepository domain.TodoRepository
}

type TodoFetchAllResponse struct {
	Todos domain.Todos `json:"items"`
}

type TodoFetchRequest struct {
	TodoID int
}

type TodoFetchResponse struct {
	Todo domain.Todo `json:"item"`
}

type TodoRegisterRequest struct {
	Todo domain.Todo
}

type TodoRegisterResponse struct {
	Todo domain.Todo `json:"item"`
}

type TodoDeleteRequest struct {
	TodoID int
}

type TodoDeleteResponse struct {
	Todo domain.Todo `json:"item"`
}

type TodoUpdateRequest struct {
	TodoID int
	Todo   domain.Todo
}

type TodoUpdateResponse struct {
	Todo domain.Todo `json:"item"`
}

func (m *TodoApplication) FetchAllFromMySQL() (*TodoFetchAllResponse, error) {
	res, err := m.TodoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &TodoFetchAllResponse{Todos: res}, nil
}

func (m *TodoApplication) FetchFromMySQL(req TodoFetchRequest) (*TodoFetchResponse, error) {
	res, err := m.TodoRepository.Find(req.TodoID)
	if err != nil {
		return nil, err
	}

	return &TodoFetchResponse{Todo: res}, nil
}

func (m *TodoApplication) RegisterToMySQL(req TodoRegisterRequest) (*TodoRegisterResponse, error) {

	res, err := m.TodoRepository.Register(req.Todo)
	if err != nil {
		return nil, err
	}

	return &TodoRegisterResponse{Todo: res}, nil
}

func (m *TodoApplication) DeleteFromMySQL(req TodoDeleteRequest) (*TodoDeleteResponse, error) {
	res, err := m.TodoRepository.Delete(req.TodoID)
	if err != nil {
		return nil, err
	}

	return &TodoDeleteResponse{Todo: res}, nil
}

func (m *TodoApplication) UpdateToMySQL(req TodoUpdateRequest) (*TodoUpdateResponse, error) {

	res, err := m.TodoRepository.Update(req.TodoID, req.Todo)
	if err != nil {
		return nil, err
	}

	return &TodoUpdateResponse{Todo: res}, nil
}
