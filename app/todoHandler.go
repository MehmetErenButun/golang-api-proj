package app

import (
	"github.com/MehmetErenButun/golang-api-proj/model"
	"github.com/MehmetErenButun/golang-api-proj/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TodoHandle struct {
	Service services.TodoService
}

func (h TodoHandle) CreateTodo(c *fiber.Ctx) error {
	var todo model.Todo

	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Service.TodoInsert(todo)

	if err != nil || result.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandle) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.GetAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}
