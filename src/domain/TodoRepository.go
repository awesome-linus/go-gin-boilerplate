package domain

type TodoRepository interface {
	FindAll() (Todos, error)
	Find(todoID int) (Todo, error)
	Register(todo Todo) (Todo, error)
	Delete(todoID int) (Todo, error)
	Update(todoID int, todo Todo) (Todo, error)
}
