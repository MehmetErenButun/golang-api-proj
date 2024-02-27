package services

import (
	"github.com/MehmetErenButun/golang-api-proj/dto"
	"github.com/MehmetErenButun/golang-api-proj/model"
	"github.com/MehmetErenButun/golang-api-proj/repository"
	"log"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}
type TodoService interface {
	TodoInsert(todo model.Todo) (dto *dto.TodoDTO, err error)
	GetAll() (todos []model.Todo, err error)
}

func (t DefaultTodoService) TodoInsert(todo model.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 3 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) GetAll() ([]model.Todo, error) {

	result, err := t.Repo.GetAll()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return result, nil
}
func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
