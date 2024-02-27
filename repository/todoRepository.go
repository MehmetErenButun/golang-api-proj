package repository

import (
	"context"
	"errors"
	"github.com/MehmetErenButun/golang-api-proj/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type TodoRepositoryDb struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo model.Todo) (bool, error)
	GetAll() ([]model.Todo, error)
}

func (t TodoRepositoryDb) Insert(todo model.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID == nil || err != nil {
		errors.New("failed add")
		return false, err
	}

	return true, nil
}

func (t TodoRepositoryDb) GetAll() ([]model.Todo, error) {
	var todo model.Todo
	var todos []model.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func NewTodoRepository(dbClient *mongo.Collection) TodoRepositoryDb {
	return TodoRepositoryDb{TodoCollection: dbClient}
}
