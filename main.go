package main

import (
	"github.com/MehmetErenButun/golang-api-proj/app"
	"github.com/MehmetErenButun/golang-api-proj/configs"
	"github.com/MehmetErenButun/golang-api-proj/repository"
	"github.com/MehmetErenButun/golang-api-proj/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDb()

	dbClient := configs.GetCollection(configs.DB, "todos")
	TodoRepositoryDb := repository.NewTodoRepository(dbClient)

	td := app.TodoHandle{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Listen(":1903")

}
